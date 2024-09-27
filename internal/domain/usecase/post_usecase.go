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

func (u *postUseCase) Broadcast(ctx context.Context, ch chan<- []byte) {
	dataCh := make(chan port.PostData)
	defer close(dataCh)

	go u.repository.ListenForChanges(ctx, dataCh)

	for {
		select {
		case <-ctx.Done():
			return
		case d := <-dataCh:
			bytes, err := u.presenter.ConvertBytes(d)
			if err != nil {
				continue
			}
			ch <- bytes
		}
	}
}
