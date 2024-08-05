package repository

import (
	"context"
	"maptalk/internal/domain/usecase/port"
	"maptalk/mock"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := mock.NewMockDataStore(ctrl)
	r := NewUserRepository(ds)
	got, err := r.FindByID("1")
	want := port.UserData{
		ID:   "1",
		Name: "John Doe",
	}
	if got != want || err != nil {
		t.Errorf("FindByID() = %v, %v, want match for %v, nil", got, err, want)
	}
}

func TestSave(t *testing.T) {
	// when
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := mock.NewMockDataStore(ctrl)
	r := NewUserRepository(ds)
	user := port.UserData{
		ID:   "1",
		Name: "John Doe",
	}
	ds.EXPECT().InsertData(gomock.Any(), gomock.Any())

	context := context.Background()
	got, err := r.Save(user, context)
	want := port.UserData{
		ID:   "1",
		Name: "John Doe",
	}
	if got != want || err != nil {
		t.Errorf("Save() = %v, %v, want match for %v, nil", got, err, want)
	}
}