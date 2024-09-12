package datastore

import (
	repo "maptalk/internal/interface/repository/port"

	"context"

	"cloud.google.com/go/firestore"
)

type Datastore struct {
	projectID string
}

func NewDataStore(projectID string) repo.DataStore {
	return &Datastore{
		projectID: projectID,
	}
}

func (ds *Datastore) GetUserData(ctx context.Context, id string) (repo.UserData, error) {
	client, err := firestore.NewClient(ctx, ds.projectID)
	if err != nil {
		return repo.UserData{}, err
	}
	defer client.Close()

	// Get data
	doc, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return repo.UserData{}, err
	}
	data := doc.Data()

	return repo.UserData{
		ID:     id,
		Name:   data["name"].(string),
		Avatar: data["avatar"].(string),
	}, nil
}

func (ds *Datastore) InsertUserData(ctx context.Context, user repo.UserInsertData) (repo.UserData, error) {
	client, err := firestore.NewClient(ctx, ds.projectID)
	if err != nil {
		return repo.UserData{}, err
	}
	defer client.Close()

	// Insert data
	res, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"name":   user.Name,
		"avatar": user.Avatar,
	})
	if err != nil {
		return repo.UserData{}, err
	}

	// Get data
	doc, err := client.Collection("users").Doc(res.ID).Get(ctx)
	if err != nil {
		return repo.UserData{}, err
	}
	data := doc.Data()

	return repo.UserData{
		ID:     res.ID,
		Name:   data["name"].(string),
		Avatar: data["avatar"].(string),
	}, nil
}

func (ds *Datastore) InsertPostData(ctx context.Context, post repo.PostInsertData) error {
	client, err := firestore.NewClient(ctx, ds.projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	// Insert data
	_, _, err = client.Collection("posts").Add(ctx, map[string]interface{}{
		"user_id":     post.UserID,
		"user_name":   post.UserName,
		"user_avatar": post.UserAvatar,
		"message":     post.Message,
		"latitude":    post.Latitude,
		"longitude":   post.Longitude,
	})
	if err != nil {
		return err
	}

	return nil
}
