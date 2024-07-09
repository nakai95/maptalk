package datastore

import (
	"context"
	"cloud.google.com/go/firestore"
	"fmt"
)

func NewDataStore() 
func createClient(ctx context.Context) *firestore.Client {
    projectID := "test-project"

    client, err := firestore.NewClient(ctx, projectID)
    if err != nil {
        fmt.Printf("error")
    }
    return client
}

func insertData(ctx context.Context, client *firestore.Client, name string) {
    _, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
        "id": "00000",
        "name": name,
    })
    if err != nil {
        fmt.Printf("error")
    }
}