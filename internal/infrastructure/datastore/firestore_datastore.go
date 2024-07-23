package datastore

import (
    "maptalk/internal/interface/repository"

	"context"
	"cloud.google.com/go/firestore"
	"fmt"
)

type Datastore struct {
    client *firestore.Client
}

func NewDataStore(ctx context.Context) (*Datastore, error) {
    projectID := "test-project"

    client, err := firestore.NewClient(ctx, projectID)
    if err != nil {
        return nil, err
    }
    return &Datastore{client: client}, nil
}

func (ds *Datastore) Close() error {
    return ds.client.Close()
}

func (ds *Datastore) InsertData(ctx context.Context, user repository.UserAccessData) {
    _, _, err := ds.client.Collection("users").Add(ctx, map[string]interface{}{
        "id": user.ID,
        "name": user.Name,
    })
    if err != nil {
        fmt.Printf("error")
    }
}