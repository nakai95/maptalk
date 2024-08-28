package port

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

// UseCase
type DraftUser struct {
	Name string
}
type UserData struct {
	ID   string
	Name string
}
type UserUseCase interface {
	GetUserByID(id string) (UserOutputData, error)
	Save(draft DraftUser, ctx context.Context) (UserOutputData, error)
}

// Presenter
type UserOutputData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserPresenter interface {
	PresentUser(user UserData) (UserOutputData, error)
}

// Repository
type UserRepository interface {
	FindByID(id string) (UserData, error)
	Save(draft DraftUser, ctx context.Context) (UserData, error)
}
