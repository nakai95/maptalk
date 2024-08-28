package usecase

import (
	"context"
	"maptalk/internal/domain/entity"
	"maptalk/internal/domain/usecase/port"
)

// UseCase
type userUseCase struct {
	presenter  port.UserPresenter
	repository port.UserRepository
}

func NewUserUseCase(presenter port.UserPresenter, repository port.UserRepository) port.UserUseCase {
	return &userUseCase{
		presenter:  presenter,
		repository: repository,
	}
}

func (u *userUseCase) GetUserByID(id string) (port.UserOutputData, error) {
	userData, err := u.repository.FindByID(id)
	if err != nil {
		return port.UserOutputData{}, err
	}
	// Convert user data to domain entity
	userEntity := entity.User(userData)

	// business logic

	// Convert domain entity to user data
	userData = port.UserData(userEntity)

	return u.presenter.PresentUser(userData)
}

func (u *userUseCase) Save(draft port.DraftUser, ctx context.Context) (port.UserOutputData, error) {
	// input -> data
	userData, err := u.repository.Save(draft, ctx)
	if err != nil {
		return port.UserOutputData{}, err
	}
	// data -> entity
	user := entity.User(userData)

	// business logic

	// entity -> data
	userData = port.UserData(user)

	// data -> output
	userOutputData, err := u.presenter.PresentUser(userData)
	if err != nil {
		return port.UserOutputData{}, err
	}
	return userOutputData, nil
}
