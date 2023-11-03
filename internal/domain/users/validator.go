package users

import (
	"context"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors"
	"strings"
)

type userRepoForValidator interface {
	ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) (bool, error)
}

type Validator struct {
	userRepo userRepoForValidator
}

type validateTask func(ctx context.Context, u User, errch chan<- string, dch chan<- struct{})

func NewValidator(repo userRepoForValidator) *Validator {
	return &Validator{userRepo: repo}
}

func (v *Validator) Validate(ctx context.Context, user User) error {
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

func (v *Validator) validateAge(_ context.Context, u User, errch chan<- string, dch chan<- struct{}) {
	if u.Age < 18 {
		errch <- onboardingerrors.ErrorAgeMinimum
	}
	dch <- struct{}{}
}

func (v *Validator) validateEmail(_ context.Context, u User, errch chan<- string, dch chan<- struct{}) {
	if u.Email == "" {
		errch <- onboardingerrors.ErrorEmailRequired
	} else if strings.Contains(u.Email, "@") == false {
		errch <- onboardingerrors.ErrorEmailFormat
	}

	dch <- struct{}{}
}

func (v *Validator) validateName(ctx context.Context, u User, errch chan<- string, dch chan<- struct{}) {
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
