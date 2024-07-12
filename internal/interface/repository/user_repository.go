package repository

import (
	"maptalk/internal/domain/usecase/port"
)

type userRepository struct{}

func NewUserRepository() port.UserDataAccess {
	return &userRepository{}
}

func (p *userRepository) FindByID(id string) (port.UserData, error) {
	// dummy data
	user := port.UserData{
		ID:   id,
		Name: "John Doe",
	}
	return user, nil
}
