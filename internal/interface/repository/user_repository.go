package repository

import (
	"context"
	"maptalk/internal/domain/usecase"
)


type userRepository struct{
    datastore DataStore
}

type DataStore interface {
    InsertData(ctx context.Context, user UserAccessData)
}

type UserAccessData struct {
    ID   string
    Name string
}

func NewUserRepository(datastore DataStore) usecase.UserDataAccess {
	return &userRepository{
        datastore: datastore,
    }
}

func (p *userRepository) FindByID(id string) (*usecase.UserData, error) {
    // dummy data
    user := &usecase.UserData{
        ID:   id,
        Name: "John Doe",
    }
    return user, nil
}

func (repo *userRepository) Save(user usecase.UserData, ctx context.Context) (*usecase.UserData, error) {
    userData := &usecase.UserData{
        ID:   user.ID,
        Name: user.Name,
    }
    userAccessData := &UserAccessData{
        ID:   userData.ID,
        Name: userData.Name,
    }
    repo.datastore.InsertData(ctx, *userAccessData)
    return userData, nil
} 