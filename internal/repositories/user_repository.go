package repositories

import (
	"boiler/internal/database"
	"boiler/internal/models"
)

// UserRepository handles user-related database operations
type UserRepository struct{}

// NewUserRepository creates a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetUser gets a user by id from the database
func (r *UserRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
