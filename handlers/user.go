package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) FindUser(c *gin.Context) {
	id := c.Param("id")

	// TODO: Handle error when user is not found
	user := h.userService.FindById(id)

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

	// TODO: Handle error when insertion fails
	h.userService.Insert(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
