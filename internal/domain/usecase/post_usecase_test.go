package usecase

import (
	"errors"
	"testing"

	"maptalk/internal/domain/entity"
	"maptalk/internal/domain/usecase/port"
	"maptalk/mock"

	"go.uber.org/mock/gomock"
)

func TestSavePost(t *testing.T) {
	// dummy data
	user := entity.User{
		ID:     "XXXXXX",
		Name:   "John Doe",
		Avatar: "/avatar/avatar1.png",
	}
	message := "Hello, World!"
	coordinate := entity.Coordinate{Latitude: 0.0, Longitude: 0.0}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock PostRepository
	r := mock.NewMockPostRepository(ctrl)
	r.EXPECT().Save(port.DraftPost{
		User:       user,
		Message:    message,
		Coordinate: coordinate,
	}, gomock.Any()).Return(nil)

	// mock PostPresenter
	p := mock.NewMockPostPresenter(ctrl)

	// create PostUseCase
	u := NewPostUseCase(p, r)

	// input data
	draft := port.DraftPost{
		User:       user,
		Message:    message,
		Coordinate: coordinate,
	}

	// when
	err := u.Save(draft, nil)

	// compare
	if err != nil {
		t.Errorf("Save() = %v, want nil", err)
	}
}

func TestSavePostFailed(t *testing.T) {
	// dummy data
	user := entity.User{
		ID:     "XXXXXX",
		Name:   "John Doe",
		Avatar: "/avatar/avatar1.png",
	}
	message := "Hello, World!"
	coordinate := entity.Coordinate{Latitude: 0.0, Longitude: 0.0}
	saveErr := errors.New("failed to save post")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock PostRepository
	r := mock.NewMockPostRepository(ctrl)
	r.EXPECT().Save(port.DraftPost{
		User:       user,
		Message:    message,
		Coordinate: coordinate,
	}, gomock.Any()).Return(saveErr)

	// mock PostPresenter
	p := mock.NewMockPostPresenter(ctrl)

	// create PostUseCase
	u := NewPostUseCase(p, r)

	// input data
	draft := port.DraftPost{
		User:       user,
		Message:    message,
		Coordinate: coordinate,
	}

	// when
	err := u.Save(draft, nil)

	// want
	want := saveErr

	// compare
	if err != want {
		t.Errorf("Save() = %v, want %v", err, want)
	}
}
