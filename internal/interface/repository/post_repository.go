package repository

import (
	"context"
	"maptalk/internal/domain/usecase/port"
	repo "maptalk/internal/interface/repository/port"
)

type PostRepository struct {
	datastore repo.DataStore
}

func NewPostRepository(datastore repo.DataStore) port.PostRepository {
	return &PostRepository{
		datastore: datastore,
	}
}

func (r *PostRepository) Save(draft port.DraftPost, ctx context.Context) error {
	// input -> data
	data := repo.PostInsertData{
		UserID:     draft.User.ID,
		UserName:   draft.User.Name,
		UserAvatar: draft.User.Avatar,
		Message:    draft.Message,
		Latitude:   draft.Coordinate.Latitude,
		Longitude:  draft.Coordinate.Longitude,
	}

	// save data
	err := r.datastore.InsertPostData(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) ListenForChanges(ctx context.Context, ch chan<- port.PostData) {
	data := make(chan repo.PostSavedData)
	defer close(data)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				r.datastore.ListenPostData(ctx, data)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case d := <-data:
			postData := r.convertDataToPostData(d)
			ch <- postData
		}
	}
}

func (r *PostRepository) convertDataToPostData(data repo.PostSavedData) port.PostData {
	postData := port.PostData{
		ID: data.ID,
		User: port.UserOutputData{
			ID:     data.UserID,
			Name:   data.UserName,
			Avatar: data.UserAvatar,
		},
		Message:   data.Message,
		CreatedAt: data.CreatedAt,
	}
	postData.Coordinate.Latitude = data.Latitude
	postData.Coordinate.Longitude = data.Longitude

	return postData
}
