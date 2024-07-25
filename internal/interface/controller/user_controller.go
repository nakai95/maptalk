package controller

import (
	"maptalk/internal/domain/usecase"
	"maptalk/internal/domain/usecase/port"
)

type UserController struct {
	userUseCase port.UserInput
}

func NewUserController(presenter port.UserOutput, repository port.UserDataAccess) *UserController {
	u := usecase.NewUserUseCase(presenter, repository)
	return &UserController{
		userUseCase: u,
	}
}

func (c *UserController) GetUserByID(id string) (port.UserOutputData, error) {
	user, err := c.userUseCase.GetUserByID(id)
	if err != nil {
		return port.UserOutputData{}, err
	}
	return user, nil
}
