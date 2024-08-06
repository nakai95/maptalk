package repositoryPort

import (
	"context"
)
//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

type UserAccessData struct {
	ID   string
	Name string
}

type DataStore interface {
    InsertData(ctx context.Context, data UserAccessData)
}
