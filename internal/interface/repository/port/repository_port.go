package repositoryPort

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

type DataStore interface {
	GetUserData(ctx context.Context, id string) (UserData, error)
	InsertUserData(ctx context.Context, data UserInsertData) (UserData, error)
	InsertPostData(ctx context.Context, data PostInsertData) error
}
