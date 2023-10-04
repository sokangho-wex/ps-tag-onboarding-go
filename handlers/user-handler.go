package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"net/http"
)

type userRepo interface {
	FindByID(id string) (models.User, error)
	AddUser(user models.User) error
}

type UserHandler struct {
	userRepo userRepo
}

func NewUserHandler(repo userRepo) *UserHandler {
	return &UserHandler{userRepo: repo}
}

func (h *UserHandler) FindUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userRepo.FindByID(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) SaveUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		err = models.BadRequestError
		_ = c.Error(err)
		return
	}

	// TODO: Add validation logic

	err := h.userRepo.AddUser(user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
