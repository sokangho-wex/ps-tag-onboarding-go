package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models/errs"
	"net/http"
)

type userRepo interface {
	FindByID(ctx context.Context, id string) (models.User, error)
	SaveUser(ctx context.Context, user models.User) error
}

type validator interface {
	Validate(ctx context.Context, user models.User) error
}

type UserHandler struct {
	userRepo  userRepo
	validator validator
}

func NewUserHandler(repo userRepo, validator validator) *UserHandler {
	return &UserHandler{userRepo: repo, validator: validator}
}

func (h *UserHandler) FindUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userRepo.FindByID(c, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) SaveUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		err = errs.NewBadRequestError()
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Saved",
	})
}