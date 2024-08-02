package controller

import (
	"context"
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

func (c *UserController) Save(name string, ctx context.Context) (port.UserOutputData, error){
	user, err := c.userUseCase.Save(name, ctx)
	if err != nil {
		return port.UserOutputData{}, err
	}
	return user, nil
}