package repository

import (
	"context"
	"errors"
	usecase "maptalk/internal/domain/usecase/port"
	"maptalk/mock"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestSavePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock DataStore
	ds := mock.NewMockDataStore(ctrl)
	ds.EXPECT().InsertPostData(gomock.Any(), gomock.Any()).Return(nil)

	// create UserRepository
	r := NewPostRepository(ds)

	// input data
	draft := usecase.DraftPost{
		Message: "Hello, World!",
	}
	draft.User.ID = "XXXXXX"
	draft.User.Name = "John Doe"
	draft.User.Avatar = "/avatar/avatar1.png"
	draft.Coordinate.Latitude = 0.0
	draft.Coordinate.Longitude = 0.0

	context := context.Background()

	// when
	err := r.Save(draft, context)

	// compare
	if err != nil {
		t.Errorf("Save() = %v, want nil", err)
	}
}

func TestSavePostFailed(t *testing.T) {
	// dummy data
	saveError := errors.New("error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock DataStore
	ds := mock.NewMockDataStore(ctrl)
	ds.EXPECT().InsertPostData(gomock.Any(), gomock.Any()).Return(saveError)

	// create UserRepository
	r := NewPostRepository(ds)

	// input data
	draft := usecase.DraftPost{
		Message: "Hello, World!",
	}
	draft.User.ID = "XXXXXX"
	draft.User.Name = "John Doe"
	draft.User.Avatar = "/avatar/avatar1.png"
	draft.Coordinate.Latitude = 0.0
	draft.Coordinate.Longitude = 0.0

	context := context.Background()

	// when
	err := r.Save(draft, context)

	// want
	want := saveError

	// compare
	if err != want {
		t.Errorf("Save() = %v, want %v", err, want)
	}
}
