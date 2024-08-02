package repository

import (
	"context"
	"maptalk/internal/domain/usecase/port"
    "maptalk/internal/interface/repository/port"
)

type UserRepository struct{
    datastore repositoryPort.DataStore
}

func NewUserRepository(datastore repositoryPort.DataStore) port.UserDataAccess {
	return &UserRepository{
        datastore: datastore,
    }
}

func (p *UserRepository) FindByID(id string) (port.UserData, error) {
	// dummy data
	user := port.UserData{
		ID:   id,
		Name: "John Doe",
	}
	return user, nil
}

func (repo *UserRepository) Save(user port.UserData, ctx context.Context) (port.UserData, error) {
    userData := port.UserData{
        ID:   user.ID,
        Name: user.Name,
    }

    data := repositoryPort.UserAccessData{
        ID:   userData.ID,
        Name: userData.Name,
    }
    repo.datastore.InsertData(ctx, data)
    return userData, nil
} 