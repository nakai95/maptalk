package repository

import (
	"maptalk/internal/domain/usecase/port"
	"testing"
)

func TestFindById(t *testing.T) {
	r := NewUserRepository()
	got, err := r.FindByID("1")
	want := port.UserData{
		ID:   "1",
		Name: "John Doe",
	}
	if got != want || err != nil {
		t.Errorf("FindByID() = %v, %v, want match for %v, nil", got, err, want)
	}
}
