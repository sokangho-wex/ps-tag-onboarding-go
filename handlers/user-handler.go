package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"net/http"
)

type userRepo interface {
	FindByID(id string) models.User
	AddUser(user models.User)
}

type UserHandler struct {
	userRepo userRepo
}

func NewUserHandler(repo userRepo) *UserHandler {
	return &UserHandler{userRepo: repo}
}

func (h *UserHandler) FindUser(c *gin.Context) {
	id := c.Param("id")

	// TODO: Handle error when user is not found
	user := h.userRepo.FindByID(id)

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) SaveUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// TODO: Add validation logic

	// TODO: Handle error when insertion fails
	h.userRepo.AddUser(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
