package port

import (
	"context"
	"maptalk/internal/domain/entity"
)

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

// UseCase
type DraftPost struct {
	User       entity.User
	Message    string
	Coordinate entity.Coordinate
}

type PostData struct {
	ID         string
	User       UserOutputData
	Message    string
	Coordinate entity.Coordinate
	CreatedAt  int64 // Unix time
}

type PostUseCase interface {
	Save(draft DraftPost, ctx context.Context) error
	Broadcast(send func([]byte), ctx context.Context) error
}

// Presenter
type PostOutputData struct {
	ID         string         `json:"id"`
	User       UserOutputData `json:"user"`
	Message    string         `json:"message"`
	Coordinate struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinate"`
	CreatedAt int64 `json:"createdAt"`
}

type PostPresenter interface {
	PresentPost(postData PostData) (PostOutputData, error)
	ConvertBytes(postData PostData) ([]byte, error)
}

// Repository
type PostRepository interface {
	Save(draft DraftPost, ctx context.Context) error
	ListenForChanges(ctx context.Context, ch chan<- PostData) error
}
