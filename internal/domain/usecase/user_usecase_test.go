package usecase

import (
	"testing"

	"maptalk/internal/domain/usecase/port"
	"maptalk/mock"

	"github.com/golang/mock/gomock"
)

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock.NewMockUserDataAccess(ctrl)
	r.EXPECT().FindByID("1").Return(port.UserData{
		ID:   "1",
		Name: "John Doe",
	}, nil)

	p := mock.NewMockUserOutput(ctrl)
	p.EXPECT().PresentUser(port.UserData{
		ID:   "1",
		Name: "John Doe",
	}).Return(port.UserOutputData{
		ID:   "1",
		Name: "John Doe",
	}, nil)

	u := NewUserUseCase(p, r)
	got, err := u.GetUserByID("1")
	want := port.UserOutputData{
		ID:   "1",
		Name: "John Doe",
	}
	if got != want || err != nil {
		t.Errorf("GetUserByID() = %v, %v, want match for %v, nil", got, err, want)
	}
}
