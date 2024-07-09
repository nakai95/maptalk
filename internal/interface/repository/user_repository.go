package repository

import (
	"context"
	"maptalk/internal/domain/usecase"

	"cloud.google.com/go/firestore"
)


type userRepository struct{
    datastore DataStore
}

type DataStore interface {
    createClient(ctx context.Context) (*firestore.Client)
    insertData(ctx context.Context, client *firestore.Client, name string)
}

func NewUserRepository() usecase.UserDataAccess {
	return &userRepository{
        datastore: datastore
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

func (repo *userRepository) save(name string) (string, error) {
    cont := context.Background()
    client := repo.datastore.createClient(cont)
    defer client.Close()
    repo.datastore.insertData(cont, client, name)
    return name, nil
} 