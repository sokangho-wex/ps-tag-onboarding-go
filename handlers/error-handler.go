package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models/errs"
	"net/http"
)

type errorResponse struct {
	Error   string   `json:"error"`
	Details []string `json:"details,omitempty"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var (
			badRequestError     *errs.BadRequestError
			userNotFoundError   *errs.NotFoundError
			userValidationError *errs.ValidationError
		)

		switch {
		case errors.As(err, &badRequestError):
			c.JSON(
				http.StatusBadRequest,
				errorResponse{Error: badRequestError.Message},
			)
		case errors.As(err, &userNotFoundError):
			c.JSON(
				http.StatusNotFound,
				errorResponse{Error: userNotFoundError.Message},
			)
		case errors.As(err, &userValidationError):
			c.JSON(
				http.StatusBadRequest,
				errorResponse{
					Error:   userValidationError.Message,
					Details: userValidationError.Details,
				},
			)
		default:
			c.JSON(
				http.StatusInternalServerError,
				errorResponse{Error: errs.ErrorUnexpected},
			)
		}
	}
}
