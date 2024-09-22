package controller

import (
	"errors"
	"testing"

	"maptalk/mock"

	"go.uber.org/mock/gomock"
)

func TestSavePost(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock PostRepository
	r := mock.NewMockPostRepository(ctrl)
	r.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

	// mock PostPresenter
	p := mock.NewMockPostPresenter(ctrl)

	// create PostUseCase
	c := NewPostController(p, r)

	// input data
	data := FormattedPost{
		Message: "Hello, World!",
	}
	data.User.ID = "XXXXXX"
	data.User.Name = "John Doe"
	data.User.Avatar = "/avatar/avatar1.png"
	data.Coordinate.Latitude = 0.0
	data.Coordinate.Longitude = 0.0

	// when
	err := c.Save(data, nil)

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

	// mock PostRepository
	r := mock.NewMockPostRepository(ctrl)
	r.EXPECT().Save(gomock.Any(), gomock.Any()).Return(saveError)

	// mock PostPresenter
	p := mock.NewMockPostPresenter(ctrl)

	// create PostUseCase
	c := NewPostController(p, r)

	// input data
	data := FormattedPost{
		Message: "Hello, World!",
	}
	data.User.ID = "XXXXXX"
	data.User.Name = "John Doe"
	data.User.Avatar = "/avatar/avatar1.png"
	data.Coordinate.Latitude = 0.0
	data.Coordinate.Longitude = 0.0

	// when
	err := c.Save(data, nil)

	// want
	want := saveError

	// compare
	if err != want {
		t.Errorf("Save() = %v, want %v", err, want)
	}
}
