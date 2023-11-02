package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models/errs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorHandler_WhenGivenAnErrorType_ReturnsCorrectStatusCodeAndMessage(t *testing.T) {
	testCases := []struct {
		name       string
		error      error
		statusCode int
		expected   models.ErrorResponse
	}{
		{
			name:       "should return 400 error response when getting BadRequestError",
			error:      errs.NewBadRequestError(),
			statusCode: http.StatusBadRequest,
			expected:   models.ErrorResponse{Error: errs.ErrorBadRequest},
		},
		{
			name:       "should return 400 error response when getting UserValidationError",
			error:      errs.NewValidationError([]string{errs.ErrorEmailFormat, errs.ErrorAgeMinimum}),
			statusCode: http.StatusBadRequest,
			expected:   models.ErrorResponse{Error: errs.ErrorValidationFailed, Details: []string{errs.ErrorEmailFormat, errs.ErrorAgeMinimum}},
		},
		{
			name:       "should return 404 error response when getting UserNotFoundError",
			error:      errs.NewNotFoundError(),
			statusCode: http.StatusNotFound,
			expected:   models.ErrorResponse{Error: errs.ErrorUserNotFound},
		},
		{
			name:       "should return 500 error response when getting UnexpectedError",
			error:      errors.New("something unexpected"),
			statusCode: http.StatusInternalServerError,
			expected:   models.ErrorResponse{Error: errs.ErrorUnexpected},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			_ = ctx.Error(tc.error)

			ErrorHandler()(ctx)
			var actualResponse models.ErrorResponse

			assert.Equal(t, tc.statusCode, response.Code)
			assert.NoError(t, json.Unmarshal(response.Body.Bytes(), &actualResponse))
			assert.Equal(t, tc.expected, actualResponse)
		})
	}
}
