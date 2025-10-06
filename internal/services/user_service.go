package services

import (
	"boiler/internal/models"
	"boiler/internal/repositories"
)

// UserService handles user-related business logic
type UserService struct {
	UserRepository *repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepository.CreateUser(user)
}

// GetUser gets a user by id
func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.UserRepository.GetUser(id)
}
