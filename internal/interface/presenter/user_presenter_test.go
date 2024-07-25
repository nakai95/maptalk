package presenter

import (
	"maptalk/internal/domain/usecase/port"
	"testing"
)

func TestPresentUser(t *testing.T) {
	user := port.UserData{
		ID:   "1",
		Name: "test",
	}
	p := NewUserPresenter()
	got, err := p.PresentUser(user)
	want := port.UserOutputData{
		ID:   "1",
		Name: "test",
	}
	if got != want || err != nil {
		t.Errorf("PresentUser() = %v, %v, want match for %v, nil", got, err, want)
	}
}
