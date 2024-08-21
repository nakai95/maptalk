package repositoryPort

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

type UserInsertData struct {
	Name string
}

type UserData struct {
	ID   string
	Name string
}

type DataStore interface {
	GetData(ctx context.Context, id string) (UserData, error)
	InsertData(ctx context.Context, data UserInsertData) (UserData, error)
}
