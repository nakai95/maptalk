package presenter

import (
	"maptalk/internal/domain/usecase/port"
)

type userPresenter struct {
}

func NewUserPresenter() port.UserOutput {
	return &userPresenter{}
}

func (p *userPresenter) PresentUser(user port.UserData) (port.UserOutputData, error) {
	userOutputData := port.UserOutputData(user)
	return userOutputData, nil
}
