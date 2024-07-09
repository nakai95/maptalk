package controller

import (
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

func (c *UserController) seave(name string) (name, error){
	user, err := c.userUseCase.save(name)
}