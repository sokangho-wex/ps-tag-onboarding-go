package validators

import (
	"context"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserValidatorMock struct {
	mock.Mock
}

func (m *UserValidatorMock) Validate(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
