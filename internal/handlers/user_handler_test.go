package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/validators"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models/errs"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserHandlerTestSuite struct {
	suite.Suite
	ctx       context.Context
	userRepo  *persistence.UserRepoMock
	validator *validators.UserValidatorMock
	service   *UserHandler
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.ctx = context.TODO()
	s.userRepo = &persistence.UserRepoMock{}
	s.validator = &validators.UserValidatorMock{}
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

	user := models.NewUser("1", "John", "Doe", "john.doe@test.com", 18)
	s.userRepo.On("FindByID", ctx, "1").Return(user, nil)

	s.service.FindUser(ctx)
	var actualUser models.User

	assert.Equal(s.T(), http.StatusOK, response.Code)
	assert.NoError(s.T(), json.Unmarshal(response.Body.Bytes(), &actualUser))
	assert.Equal(s.T(), user, actualUser)
}

func (s *UserHandlerTestSuite) TestFindUser_WhenUserNotFound() {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request = httptest.NewRequest("GET", "/find", nil)
	ctx.Params = gin.Params{{Key: "id", Value: "1"}}

	err := errs.NewNotFoundError()
	s.userRepo.On("FindByID", ctx, "1").Return(models.User{}, err)

	s.service.FindUser(ctx)

	assert.Equal(s.T(), ctx.Errors.Last().Err, err)
}
