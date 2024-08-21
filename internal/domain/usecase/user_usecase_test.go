package usecase

import (
	"testing"

	"maptalk/internal/domain/usecase/port"
	"maptalk/mock"

	"go.uber.org/mock/gomock"
)

func TestGetUserByID(t *testing.T) {
	// dummy data
	id := "XXXXXX"
	name := "John Doe"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock UserRepository
	r := mock.NewMockUserRepository(ctrl)
	r.EXPECT().FindByID(id).Return(port.UserData{
		ID:   id,
		Name: name,
	}, nil)

	// mock UserPresenter
	p := mock.NewMockUserPresenter(ctrl)
	p.EXPECT().PresentUser(port.UserData{
		ID:   id,
		Name: name,
	}).Return(port.UserOutputData{
		ID:   id,
		Name: name,
	}, nil)

	// create UserUseCase
	u := NewUserUseCase(p, r)

	// when
	got, err := u.GetUserByID(id)

	// then
	want := port.UserOutputData{
		ID:   id,
		Name: name,
	}

	// compare
	if got != want || err != nil {
		t.Errorf("GetUserByID() = %v, %v, want match for %v, nil", got, err, want)
	}
}

func TestSave(t *testing.T) {
	// dummy data
	id := "XXXXXX"
	name := "John Doe"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock UserRepository
	r := mock.NewMockUserRepository(ctrl)
	r.EXPECT().Save(port.DraftUser{
		Name: name,
	}, gomock.Any()).Return(port.UserData{
		ID:   id,
		Name: name,
	}, nil)

	// mock UserPresenter
	p := mock.NewMockUserPresenter(ctrl)
	p.EXPECT().PresentUser(port.UserData{
		ID:   id,
		Name: name,
	}).Return(port.UserOutputData{
		ID:   id,
		Name: name,
	}, nil)

	// input data
	draft := port.DraftUser{
		Name: name,
	}

	// create UserUseCase
	u := NewUserUseCase(p, r)

	// when
	got, err := u.Save(draft, nil)

	// then
	want := port.UserOutputData{
		ID:   id,
		Name: name,
	}

	// compare
	if got != want || err != nil {
		t.Errorf("Save() = %v, %v, want match for %v, nil", got, err, want)
	}
}
