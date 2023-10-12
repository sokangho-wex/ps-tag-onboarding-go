package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"net/http"
)

func FindUser(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func SaveUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
