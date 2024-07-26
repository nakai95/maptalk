package controller

import (
	"context"
	"maptalk/internal/domain/usecase"
)

type UserController struct {
	userUseCase usecase.UserInputPort
}

func NewUserController(presenter usecase.UserOutputPort, repository usecase.UserDataAccess) *UserController {
	u := usecase.NewUserUseCase(presenter, repository)
	return &UserController{
		userUseCase: u,
	}
}

func (c *UserController) GetUserByID(id string) (*usecase.UserOutputData, error) {
	user, err := c.userUseCase.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *UserController) Save(name string, ctx context.Context) (*usecase.UserOutputData, error){
	user, err := c.userUseCase.Save(name, ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}