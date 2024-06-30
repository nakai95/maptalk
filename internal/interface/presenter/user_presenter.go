package presenter

import (
	"maptalk/internal/domain/usecase"
)

type userPresenter struct {
}

func NewUserPresenter() usecase.UserOutputPort {
	return &userPresenter{}
}

func (p *userPresenter) PresentUser(user *usecase.UserData) (*usecase.UserOutputData, error)  {
	userOutputData := &usecase.UserOutputData{
		ID:   user.ID,
		Name: user.Name,
	}
	return userOutputData, nil
}