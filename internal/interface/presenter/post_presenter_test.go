package presenter

import (
	"maptalk/internal/domain/usecase/port"
	"testing"
)

func TestPresentPost(t *testing.T) {
	// input data
	post := port.PostData{
		ID:        "1",
		Message:   "Hello, World!",
		CreatedAt: 1600000000, // 2020-09-13 00:00:00
	}
	post.User.ID = "XXXX"
	post.User.Name = "John Doe"
	post.User.Avatar = "/avatar/avatar1.png"
	post.Coordinate.Latitude = 0.0
	post.Coordinate.Longitude = 0.0

	p := NewPostPresenter()

	got, err := p.PresentPost(post)

	want := port.PostOutputData{
		ID:        "1",
		Message:   "Hello, World!",
		CreatedAt: 1600000000,
	}
	want.User.ID = "XXXX"
	want.User.Name = "John Doe"
	want.User.Avatar = "/avatar/avatar1.png"
	want.Coordinate.Latitude = 0.0
	want.Coordinate.Longitude = 0.0

	if got != want || err != nil {
		t.Errorf("PresentPost() = %v, %v, want match for %v, nil", got, err, want)
	}
}

func TestConvertBytes(t *testing.T) {
	// input data
	post := port.PostData{
		ID:        "1",
		Message:   "Hello, World!",
		CreatedAt: 1600000000, // 2020-09-13 00:00:00
	}
	post.User.ID = "XXXX"
	post.User.Name = "John Doe"
	post.User.Avatar = "/avatar/avatar1.png"
	post.Coordinate.Latitude = 0.0
	post.Coordinate.Longitude = 0.0

	p := NewPostPresenter()

	got, err := p.ConvertBytes(post)

	want := []byte(`{"id":"1","user":{"id":"XXXX","name":"John Doe","avatar":"/avatar/avatar1.png"},"message":"Hello, World!","coordinate":{"latitude":0,"longitude":0},"createdAt":1600000000}`)

	if string(got) != string(want) || err != nil {
		t.Errorf("ConvertBytes() = %v, %v, want match for %v, nil", string(got), err, string(want))
	}
}
