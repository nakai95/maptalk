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
