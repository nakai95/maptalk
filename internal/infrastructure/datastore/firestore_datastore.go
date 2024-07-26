package datastore

import (
    "maptalk/internal/interface/repository"

	"context"
	"cloud.google.com/go/firestore"
	"fmt"
)

type Datastore struct {
    projectID string
}

func NewDataStore(projectID string) (*Datastore, error) {
    return &Datastore{
        projectID: projectID,
    }, nil 
}

func (ds *Datastore) InsertData(ctx context.Context, user repository.UserAccessData) {
    client, err := firestore.NewClient(ctx, ds.projectID)
    if err != nil {
        fmt.Print("error")
    }

    defer client.Close()

    _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
        "id": user.ID,
        "name": user.Name,
    })
    if err != nil {
        fmt.Printf("error")
    }
}