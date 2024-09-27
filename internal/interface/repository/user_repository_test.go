package repository

import (
	"context"
	usecase "maptalk/internal/domain/usecase/port"
	repository "maptalk/internal/interface/repository/port"
	"maptalk/mock"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestFindById(t *testing.T) {
	// dummy data
	id := "XXXXX"
	name := "John Doe"
	avatar := "/avatar/avatar1.png"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock DataStore
	ds := mock.NewMockDataStore(ctrl)
	ds.EXPECT().GetUserData(gomock.Any(), id).Return(repository.UserData{
		ID:     id,
		Name:   name,
		Avatar: avatar,
	}, nil)

	// create UserRepository
	r := NewUserRepository(ds)

	// when
	got, err := r.FindByID(id)

	// then
	want := usecase.UserData{
		ID:     id,
		Name:   name,
		Avatar: avatar,
	}

	// compare
	if got != want || err != nil {
		t.Errorf("FindByID() = %v, %v, want match for %v, nil", got, err, want)
	}
}

func TestSaveUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock DataStore
	ds := mock.NewMockDataStore(ctrl)
	ds.EXPECT().InsertUserData(gomock.Any(), gomock.Any()).Return(repository.UserData{
		ID:     "1",
		Name:   "John Doe",
		Avatar: "/avatar/avatar1.png",
	}, nil)

	// create UserRepository
	r := NewUserRepository(ds)

	// input data
	user := usecase.DraftUser{
		Name:   "John Doe",
		Avatar: "/avatar/avatar1.png",
	}
	context := context.Background()

	// when
	got, err := r.Save(user, context)

	// then
	want := usecase.UserData{
		ID:     "1",
		Name:   "John Doe",
		Avatar: "/avatar/avatar1.png",
	}

	// compare
	if got != want || err != nil {
		t.Errorf("Save() = %v, %v, want match for %v, nil", got, err, want)
	}
}
