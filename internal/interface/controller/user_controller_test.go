package controller

import (
	"maptalk/internal/domain/usecase/port"
	"testing"

	"maptalk/mock"

	"go.uber.org/mock/gomock"
)

func TestGetUserByID(t *testing.T) {
	// dummy data
	id := "XXXXX"
	name := "John Doe"
	avatar := "/avatar/avatar1.png"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock UserPresenter
	r := mock.NewMockUserRepository(ctrl)
	r.EXPECT().FindByID(id).Return(port.UserData{
		ID:     id,
		Name:   name,
		Avatar: avatar,
	}, nil)

	// mock UserPresenter
	p := mock.NewMockUserPresenter(ctrl)
	p.EXPECT().PresentUser(port.UserData{
		ID:     id,
		Name:   name,
		Avatar: avatar,
	}).Return(port.UserOutputData{
		ID:     id,
		Name:   name,
		Avatar: avatar,
	}, nil)

	// create UserController
	c := NewUserController(p, r)

	// when
	got, err := c.GetUserByID(id)

	// then
	want := port.UserOutputData{
		ID:     id,
		Name:   name,
		Avatar: avatar,
	}

	// compare
	if got != want || err != nil {
		t.Errorf("GetUserByID() = %v, %v, want match for %v, nil", got, err, want)
	}
}
