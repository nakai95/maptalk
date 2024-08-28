package controller

import (
	"context"
	"maptalk/internal/domain/usecase"
	"maptalk/internal/domain/usecase/port"
)

type UserInputData struct {
	Name string `json:"name"`
}

type userController struct {
	userUseCase port.UserUseCase
}

type UserController interface {
	GetUserByID(id string) (port.UserOutputData, error)
	Save(input UserInputData, ctx context.Context) (port.UserOutputData, error)
}

func NewUserController(presenter port.UserPresenter, repository port.UserRepository) UserController {
	u := usecase.NewUserUseCase(presenter, repository)
	return &userController{
		userUseCase: u,
	}
}

func (c *userController) GetUserByID(id string) (port.UserOutputData, error) {
	user, err := c.userUseCase.GetUserByID(id)
	if err != nil {
		return port.UserOutputData{}, err
	}
	return user, nil
}

func (c *userController) Save(input UserInputData, ctx context.Context) (port.UserOutputData, error) {
	draft := port.DraftUser(input)
	user, err := c.userUseCase.Save(draft, ctx)
	if err != nil {
		return port.UserOutputData{}, err
	}
	return user, nil
}
