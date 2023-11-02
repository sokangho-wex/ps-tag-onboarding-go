package httphandler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var (
			badRequestError     *onboardingerrors.BadRequestError
			userNotFoundError   *onboardingerrors.NotFoundError
			userValidationError *onboardingerrors.ValidationError
		)

		switch {
		case errors.As(err, &badRequestError):
			c.JSON(
				http.StatusBadRequest,
				errorResponse{Error: badRequestError.Message},
			)
		case errors.As(err, &userValidationError):
			c.JSON(
				http.StatusBadRequest,
				errorResponse{
					Error:   userValidationError.Message,
					Details: userValidationError.Details,
				},
			)
		case errors.As(err, &userNotFoundError):
			c.JSON(
				http.StatusNotFound,
				errorResponse{Error: userNotFoundError.Message},
			)
		default:
			c.JSON(
				http.StatusInternalServerError,
				errorResponse{Error: onboardingerrors.ErrorUnexpected},
			)
		}
	}
}
