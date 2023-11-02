package onboardingerrors

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
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
		expected   errorResponse
	}{
		{
			name:       "should return 400 error response when getting BadRequestError",
			error:      NewBadRequestError(),
			statusCode: http.StatusBadRequest,
			expected:   errorResponse{Error: ErrorBadRequest},
		},
		{
			name:       "should return 400 error response when getting UserValidationError",
			error:      NewValidationError([]string{ErrorEmailFormat, ErrorAgeMinimum}),
			statusCode: http.StatusBadRequest,
			expected:   errorResponse{Error: ErrorValidationFailed, Details: []string{ErrorEmailFormat, ErrorAgeMinimum}},
		},
		{
			name:       "should return 404 error response when getting UserNotFoundError",
			error:      NewNotFoundError(),
			statusCode: http.StatusNotFound,
			expected:   errorResponse{Error: ErrorUserNotFound},
		},
		{
			name:       "should return 500 error response when getting UnexpectedError",
			error:      errors.New("something unexpected"),
			statusCode: http.StatusInternalServerError,
			expected:   errorResponse{Error: ErrorUnexpected},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			_ = ctx.Error(tc.error)

			ErrorHandler()(ctx)
			var actualResponse errorResponse

			assert.Equal(t, tc.statusCode, response.Code)
			assert.NoError(t, json.Unmarshal(response.Body.Bytes(), &actualResponse))
			assert.Equal(t, tc.expected, actualResponse)
		})
	}
}
