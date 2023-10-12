package persistence

import (
	"context"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) FindByID(ctx context.Context, id string) (models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepoMock) SaveUser(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepoMock) ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) (bool, error) {
	args := m.Called(ctx, firstName, lastName)
	return args.Bool(0), args.Error(1)
}
