package repository

import (
	"context"
	"maptalk/internal/domain/usecase/port"
	repo "maptalk/internal/interface/repository/port"
)

type UserRepository struct {
	datastore repo.DataStore
}

func NewUserRepository(datastore repo.DataStore) port.UserRepository {
	return &UserRepository{
		datastore: datastore,
	}
}

func (r *UserRepository) FindByID(id string) (port.UserData, error) {
	// id -> data
	data, err := r.datastore.GetData(context.Background(), id)
	if err != nil {
		return port.UserData{}, err
	}
	return port.UserData(data), nil
}

func (r *UserRepository) Save(draft port.DraftUser, ctx context.Context) (port.UserData, error) {
	// input -> data
	data := repo.UserInsertData{
		Name:   draft.Name,
		Avatar: draft.Avatar,
	}

	// save data
	userData, err := r.datastore.InsertData(ctx, data)
	if err != nil {
		return port.UserData{}, err
	}

	return port.UserData(userData), nil
}
