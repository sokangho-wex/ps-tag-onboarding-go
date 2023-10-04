package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		switch {
		case errors.Is(err, models.BadRequestError):
			c.JSON(http.StatusBadRequest, models.BadRequestError.Error())
		case errors.Is(err, models.UserNotFoundError):
			c.JSON(http.StatusNotFound, models.UserNotFoundError.Error())
		default:
			c.JSON(http.StatusInternalServerError, models.UnexpectedError.Error())
		}
	}
}
