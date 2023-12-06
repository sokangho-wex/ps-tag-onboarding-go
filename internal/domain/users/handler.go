package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors"
	"net/http"
)

type userRepoForHandler interface {
	FindByID(ctx context.Context, id string) (User, error)
	SaveUser(ctx context.Context, user User) error
}

type validator interface {
	Validate(ctx context.Context, user User) error
}

type Handler struct {
	userRepo  userRepoForHandler
	validator validator
}

func NewHandler(repo userRepoForHandler, validator validator) *Handler {
	return &Handler{userRepo: repo, validator: validator}
}

func (h *Handler) FindUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userRepo.FindByID(c, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) SaveUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		err = onboardingerrors.NewBadRequestError()
		_ = c.Error(err)
		return
	}

	if err := h.validator.Validate(c, user); err != nil {
		_ = c.Error(err)
		return
	}

	if err := h.userRepo.SaveUser(c, user); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, user)
}
