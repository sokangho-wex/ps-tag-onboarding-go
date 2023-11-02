package validators

import (
	"context"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/onboardingerrors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/users"
	"strings"
)

type userRepo interface {
	ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) (bool, error)
}

type UserValidator struct {
	userRepo userRepo
}

type validateTask func(ctx context.Context, u users.User, errch chan<- string, dch chan<- struct{})

func NewUserValidator(repo userRepo) *UserValidator {
	return &UserValidator{userRepo: repo}
}

func (v *UserValidator) Validate(ctx context.Context, user users.User) error {
	var errorDetails []string

	errch := make(chan string)
	dch := make(chan struct{})
	validateTasks := []validateTask{v.validateName, v.validateEmail, v.validateAge}

	for _, task := range validateTasks {
		go task(ctx, user, errch, dch)
	}

	dCounter := 0

done:
	for {
		select {
		case err := <-errch:
			errorDetails = append(errorDetails, err)
		case <-dch:
			dCounter++
			if dCounter == len(validateTasks) {
				break done
			}
		}
	}

	if len(errorDetails) > 0 {
		return onboardingerrors.NewValidationError(errorDetails)
	}

	return nil
}

func (v *UserValidator) validateAge(_ context.Context, u users.User, errch chan<- string, dch chan<- struct{}) {
	if u.Age < 18 {
		errch <- onboardingerrors.ErrorAgeMinimum
	}
	dch <- struct{}{}
}

func (v *UserValidator) validateEmail(_ context.Context, u users.User, errch chan<- string, dch chan<- struct{}) {
	if u.Email == "" {
		errch <- onboardingerrors.ErrorEmailRequired
	} else if strings.Contains(u.Email, "@") == false {
		errch <- onboardingerrors.ErrorEmailFormat
	}

	dch <- struct{}{}
}

func (v *UserValidator) validateName(ctx context.Context, u users.User, errch chan<- string, dch chan<- struct{}) {
	if u.FirstName == "" || u.LastName == "" {
		errch <- onboardingerrors.ErrorNameRequired
	} else {
		exist, err := v.userRepo.ExistsByFirstNameAndLastName(ctx, u.FirstName, u.LastName)
		if err != nil {
			errch <- onboardingerrors.ErrorUnexpected
		} else if exist {
			errch <- onboardingerrors.ErrorNameUnique
		}
	}

	dch <- struct{}{}
}
