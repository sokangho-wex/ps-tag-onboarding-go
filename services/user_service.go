package services

import (
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/persistence"
)

type UserService struct {
	// TODO: Need to use interface instead
	userRepo *persistence.UserRepository
}

func NewUserService(userRepo *persistence.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) FindById(id string) models.User {
	return s.userRepo.FindById(id)
}

func (s *UserService) Insert(user models.User) {
	s.userRepo.Insert(user)
}
