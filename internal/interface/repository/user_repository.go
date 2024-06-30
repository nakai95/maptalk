package repository

import (
	"maptalk/internal/domain/usecase"
)


type userRepository struct{}

func NewUserRepository() usecase.UserDataAccess {
	return &userRepository{}
}

func (p *userRepository) FindByID(id string) (*usecase.UserData, error) {
    // dummy data
    user := &usecase.UserData{
        ID:   id,
        Name: "John Doe",
    }
    return user, nil
}