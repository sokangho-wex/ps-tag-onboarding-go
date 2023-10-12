package validators

import (
	"context"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models/errs"
	"strings"
)

type userRepo interface {
	ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) (bool, error)
}

type UserValidator struct {
	userRepo userRepo
}

type validateTask func(ctx context.Context, u models.User, errch chan<- string, dch chan<- struct{})

func NewUserValidator(repo userRepo) *UserValidator {
	return &UserValidator{userRepo: repo}
}

func (v *UserValidator) Validate(ctx context.Context, user models.User) error {
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
		return errs.NewValidationError(errorDetails)
	}

	return nil
}

func (v *UserValidator) validateAge(_ context.Context, u models.User, errch chan<- string, dch chan<- struct{}) {
	if u.Age < 18 {
		errch <- errs.ErrorAgeMinimum
	}
	dch <- struct{}{}
}

func (v *UserValidator) validateEmail(_ context.Context, u models.User, errch chan<- string, dch chan<- struct{}) {
	if u.Email == "" {
		errch <- errs.ErrorEmailRequired
	} else if strings.Contains(u.Email, "@") == false {
		errch <- errs.ErrorEmailFormat
	}

	dch <- struct{}{}
}

func (v *UserValidator) validateName(ctx context.Context, u models.User, errch chan<- string, dch chan<- struct{}) {
	if u.FirstName == "" || u.LastName == "" {
		errch <- errs.ErrorNameRequired
	} else {
		exist, err := v.userRepo.ExistsByFirstNameAndLastName(ctx, u.FirstName, u.LastName)
		if err != nil {
			errch <- errs.ErrorUnexpected
		} else if exist {
			errch <- errs.ErrorNameUnique
		}
	}

	dch <- struct{}{}
}
