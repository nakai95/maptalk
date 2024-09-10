package port

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

// UseCase
type DraftUser struct {
	Name   string
	Avatar string
}
type UserData struct {
	ID     string
	Name   string
	Avatar string
}
type UserUseCase interface {
	GetUserByID(id string) (UserOutputData, error)
	Save(draft DraftUser, ctx context.Context) (UserOutputData, error)
}

// Presenter
type UserOutputData struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UserPresenter interface {
	PresentUser(user UserData) (UserOutputData, error)
}

// Repository
type UserRepository interface {
	FindByID(id string) (UserData, error)
	Save(draft DraftUser, ctx context.Context) (UserData, error)
}
