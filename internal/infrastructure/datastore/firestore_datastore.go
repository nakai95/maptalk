package datastore

import (
	repo "maptalk/internal/interface/repository/port"
	"time"

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
		"created_at":  time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (ds *Datastore) PostDataListener(ctx context.Context) (chan repo.PostData, error) {
	ch := make(chan repo.PostData)

	client, err := firestore.NewClient(ctx, ds.projectID)
	if err != nil {
		return ch, err
	}

	go func() {
		defer close(ch)

		iter := client.Collection("posts").Where("created_at", ">", time.Now().Unix()).Snapshots(ctx)
		defer iter.Stop()

		for {
			snap, err := iter.Next()
			if err != nil {
				return
			}

			for _, change := range snap.Changes {
				switch change.Kind {
				case firestore.DocumentAdded, firestore.DocumentModified:
					data := change.Doc.Data()
					post := repo.PostData{
						ID:         change.Doc.Ref.ID,
						UserID:     data["user_id"].(string),
						UserName:   data["user_name"].(string),
						UserAvatar: data["user_avatar"].(string),
						Message:    data["message"].(string),
						Latitude:   data["latitude"].(float64),
						Longitude:  data["longitude"].(float64),
						CreatedAt:  data["created_at"].(int64),
					}
					ch <- post
				}
			}
		}
	}()

	return ch, nil
}
