package users

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type UserValidatorMock struct {
	mock.Mock
}

func (m *UserValidatorMock) Validate(ctx context.Context, user User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
