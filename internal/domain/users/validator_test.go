package users

import (
	"context"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserValidatorTestSuite struct {
	suite.Suite
	ctx      context.Context
	userRepo *UserRepoMock
	service  *Validator
}

func (s *UserValidatorTestSuite) SetupTest() {
	s.ctx = context.TODO()
	s.userRepo = &UserRepoMock{}
	s.service = NewValidator(s.userRepo)
}

func TestUserValidator(t *testing.T) {
	suite.Run(t, new(UserValidatorTestSuite))
}

func (s *UserValidatorTestSuite) TestValidate_WhenNameIsUnique() {
	testCases := []struct {
		name     string
		input    User
		expected error
	}{
		{
			name:     "should not return error when user is valid",
			input:    New("1", "John", "Doe", "john.doe@test.com", 18),
			expected: nil,
		},
		{
			name:     "should return error when age is invalid",
			input:    New("1", "John", "Doe", "john.doe@test.com", 17),
			expected: onboardingerrors.NewValidationError([]string{onboardingerrors.ErrorAgeMinimum}),
		},
		{
			name:     "should return error when email is invalid",
			input:    New("1", "John", "Doe", "", 25),
			expected: onboardingerrors.NewValidationError([]string{onboardingerrors.ErrorEmailRequired}),
		},
		{
			name:     "should return error when name is invalid",
			input:    New("1", "qwe", "", "john.doe@test.com", 25),
			expected: onboardingerrors.NewValidationError([]string{onboardingerrors.ErrorNameRequired}),
		},
		{
			name:     "should return errors when multiple fields are invalid",
			input:    New("1", "", "Doe", "john.doe-test.com", 17),
			expected: onboardingerrors.NewValidationError([]string{onboardingerrors.ErrorNameRequired, onboardingerrors.ErrorAgeMinimum, onboardingerrors.ErrorEmailFormat}),
		},
	}

	for _, tc := range testCases {
		s.userRepo.
			On("ExistsByFirstNameAndLastName", s.ctx, tc.input.FirstName, tc.input.LastName).
			Return(false, nil)

		s.Run(tc.name, func() {
			err := s.service.Validate(context.TODO(), tc.input)

			if err != nil {
				assert.Equal(s.T(), tc.expected.Error(), err.Error())
				assert.ElementsMatch(s.T(), tc.expected.(*onboardingerrors.ValidationError).Details, err.(*onboardingerrors.ValidationError).Details)
			} else {
				assert.Equal(s.T(), tc.expected, err)
			}
		})
	}
}

func (s *UserValidatorTestSuite) TestValidate_WhenNameIsNotUnique() {
	user := New("1", "John", "Doe", "john.doe@test.com", 18)

	s.userRepo.
		On("ExistsByFirstNameAndLastName", s.ctx, user.FirstName, user.LastName).
		Return(true, nil)

	err := s.service.Validate(context.TODO(), user)
	expectedErr := onboardingerrors.NewValidationError([]string{onboardingerrors.ErrorNameUnique})

	assert.Equal(s.T(), expectedErr, err)
}
