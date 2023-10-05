package validators

import (
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models/errs"
	"strings"
)

type userRepo interface {
	ExistsByFirstNameAndLastName(firstName string, lastName string) bool
}

type UserValidator struct {
	userRepo userRepo
}

func NewUserValidator(repo userRepo) *UserValidator {
	return &UserValidator{userRepo: repo}
}

func (v *UserValidator) Validate(user models.User) error {
	var errorDetails []string

	// TODO: Refactor to perform validation in goroutines for performance improvement
	if success, errorMessage := v.validateName(user); success == false {
		errorDetails = append(errorDetails, errorMessage)
	}

	if success, errorMessage := v.validateEmail(user); success == false {
		errorDetails = append(errorDetails, errorMessage)
	}

	if success, errorMessage := v.validateAge(user); success == false {
		errorDetails = append(errorDetails, errorMessage)
	}

	if len(errorDetails) > 0 {
		return errs.NewValidationError(errorDetails)
	}

	return nil
}

func (v *UserValidator) validateAge(u models.User) (bool, string) {
	if u.Age < 18 {
		return false, errs.ErrorAgeMinimum
	}
	return true, ""
}

func (v *UserValidator) validateEmail(u models.User) (bool, string) {
	if u.Email == "" {
		return false, errs.ErrorEmailRequired
	}

	if strings.Contains(u.Email, "@") == false {
		return false, errs.ErrorEmailFormat
	}

	return true, ""
}

func (v *UserValidator) validateName(u models.User) (bool, string) {
	if u.FirstName == "" || u.LastName == "" {
		return false, errs.ErrorNameRequired
	}

	if v.userRepo.ExistsByFirstNameAndLastName(u.FirstName, u.LastName) {
		return false, errs.ErrorNameUnique
	}

	return true, ""
}
