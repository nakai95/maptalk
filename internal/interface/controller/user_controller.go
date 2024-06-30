package controller

import (
	"maptalk/internal/domain/usecase"
)

type UserController struct {
	userUseCase usecase.UserInputPort
}

func NewUserController(u usecase.UserInputPort) *UserController {
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
