package controller

import (
	"context"
	"maptalk/internal/domain/usecase"
	"maptalk/internal/domain/usecase/port"
)

type PostInputData struct {
	UserId     string `json:"userId"`
	Message    string `json:"message"`
	Coordinate struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinate"`
}

type FormattedPost struct {
	User struct {
		ID     string
		Name   string
		Avatar string
	}
	Message    string
	Coordinate struct {
		Latitude  float64
		Longitude float64
	}
}

type postController struct {
	postUseCase port.PostUseCase
}

type PostController interface {
	Save(data FormattedPost, ctx context.Context) error
	Broadcast(send func([]byte), ctx context.Context) error
}

func NewPostController(presenter port.PostPresenter, repository port.PostRepository) PostController {
	u := usecase.NewPostUseCase(presenter, repository)
	return &postController{
		postUseCase: u,
	}
}

func (c *postController) Save(data FormattedPost, ctx context.Context) error {
	draft := port.DraftPost{
		User:       data.User,
		Message:    data.Message,
		Coordinate: data.Coordinate}

	err := c.postUseCase.Save(draft, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *postController) Broadcast(send func([]byte), ctx context.Context) error {
	err := c.postUseCase.Broadcast(send, ctx)
	if err != nil {
		return err
	}
	return nil
}
