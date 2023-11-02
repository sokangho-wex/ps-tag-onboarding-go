package users

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) FindByID(ctx context.Context, id string) (User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(User), args.Error(1)
}

func (m *UserRepoMock) SaveUser(ctx context.Context, user User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepoMock) ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) (bool, error) {
	args := m.Called(ctx, firstName, lastName)
	return args.Bool(0), args.Error(1)
}
