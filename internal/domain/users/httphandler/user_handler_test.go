package httphandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/users"
	validator2 "github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/users/validator"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type UserHandlerTestSuite struct {
	suite.Suite
	ctx       context.Context
	userRepo  *persistence.UserRepoMock
	validator *validator2.UserValidatorMock
	service   *UserHandler
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.ctx = context.TODO()
	s.userRepo = &persistence.UserRepoMock{}
	s.validator = &validator2.UserValidatorMock{}
	s.service = NewUserHandler(s.userRepo, s.validator)
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) TestFindUser_WhenUserFound() {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request = httptest.NewRequest("GET", "/find", nil)
	ctx.Params = gin.Params{{Key: "id", Value: "1"}}

	user := users.NewUser("1", "John", "Doe", "john.doe@test.com", 18)
	s.userRepo.On("FindByID", ctx, "1").Return(user, nil)

	s.service.FindUser(ctx)
	var actualUser users.User

	assert.Equal(s.T(), http.StatusOK, response.Code)
	assert.NoError(s.T(), json.Unmarshal(response.Body.Bytes(), &actualUser))
	assert.Equal(s.T(), user, actualUser)
}

func (s *UserHandlerTestSuite) TestFindUser_WhenUserNotFound() {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request = httptest.NewRequest("GET", "/find", nil)
	ctx.Params = gin.Params{{Key: "id", Value: "1"}}

	err := onboardingerrors.NewNotFoundError()
	s.userRepo.On("FindByID", ctx, "1").Return(users.User{}, err)

	s.service.FindUser(ctx)

	assert.Equal(s.T(), ctx.Errors.Last().Err, err)
}

func (s *UserHandlerTestSuite) TestSaveUser_WhenSuccessful() {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	user := users.NewUser("1", "John", "Doe", "john.doe@example.com", 18)
	body, _ := json.Marshal(user)
	ctx.Request = httptest.NewRequest("POST", "/save", strings.NewReader(string(body)))

	s.validator.On("Validate", ctx, user).Return(nil)
	s.userRepo.On("SaveUser", ctx, user).Return(nil)

	s.service.SaveUser(ctx)
	expectedResponseBody, _ := json.Marshal(user)

	assert.Equal(s.T(), http.StatusCreated, response.Code)
	assert.Equal(s.T(), expectedResponseBody, response.Body.Bytes())
}

func (s *UserHandlerTestSuite) TestSaveUser_WhenRequestBodyIsInvalid() {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	body := `{"id":"1","first_name123":"John","last_name213":"Doe","email123":"",age:""`
	ctx.Request = httptest.NewRequest("POST", "/save", strings.NewReader(body))

	err := onboardingerrors.NewBadRequestError()
	s.validator.On("Validate", ctx, users.User{}).Return(nil)
	s.userRepo.On("SaveUser", ctx, users.User{}).Return(nil)

	s.service.SaveUser(ctx)

	assert.Equal(s.T(), ctx.Errors.Last().Err, err)
}

func (s *UserHandlerTestSuite) TestSaveUser_WhenSomethingFails() {
	testCases := []struct {
		name           string
		validatorError error
		repoError      error
		expected       error
	}{
		{
			name:           "should append ValidationError to the context when validator fails",
			validatorError: onboardingerrors.NewValidationError([]string{onboardingerrors.ErrorEmailFormat, onboardingerrors.ErrorAgeMinimum}),
			repoError:      nil,
			expected:       onboardingerrors.NewValidationError([]string{onboardingerrors.ErrorEmailFormat, onboardingerrors.ErrorAgeMinimum}),
		},
		{
			name:           "should append UnexpectedError to the context when repo fails",
			validatorError: nil,
			repoError:      onboardingerrors.NewUnexpectedError(errors.New("something is wrong")),
			expected:       onboardingerrors.NewUnexpectedError(errors.New("something is wrong")),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			user := users.NewUser("1", "John", "Doe", "john.doe@example.com", 18)
			body, _ := json.Marshal(user)
			ctx.Request = httptest.NewRequest("POST", "/save", strings.NewReader(string(body)))

			s.validator.On("Validate", ctx, user).Return(tc.validatorError)
			s.userRepo.On("SaveUser", ctx, user).Return(tc.repoError)

			s.service.SaveUser(ctx)

			assert.Equal(s.T(), ctx.Errors.Last().Err, tc.expected)
		})
	}
}
