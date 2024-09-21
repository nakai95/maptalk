package usecase

import (
	"context"
	"maptalk/internal/domain/usecase/port"
)

// UseCase
type postUseCase struct {
	presenter  port.PostPresenter
	repository port.PostRepository
}

func NewPostUseCase(presenter port.PostPresenter, repository port.PostRepository) port.PostUseCase {
	return &postUseCase{
		presenter:  presenter,
		repository: repository,
	}
}

func (u *postUseCase) Save(draft port.DraftPost, ctx context.Context) error {
	// input -> data
	err := u.repository.Save(draft, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *postUseCase) Broadcast(send func([]byte), ctx context.Context) error {
	ch := make(chan port.PostData)
	err := u.repository.ListenForChanges(ctx, ch)
	if err != nil {
		return err
	}
	for post := range ch {
		bytes, err := u.presenter.ConvertBytes(post)
		if err != nil {
			return err
		}
		send(bytes)
	}
	return nil
}
