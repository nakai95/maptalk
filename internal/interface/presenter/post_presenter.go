package presenter

import (
	"maptalk/internal/domain/usecase/port"
)

type postPresenter struct {
}

func NewPostPresenter() port.PostPresenter {
	return &postPresenter{}
}

func (p *postPresenter) PresentPost(post port.PostData) (port.PostOutputData, error) {
	postOutputData := port.PostOutputData{}
	postOutputData.ID = post.ID
	postOutputData.User = port.UserOutputData(post.User)
	postOutputData.Message = post.Message
	postOutputData.Coordinate.Latitude = post.Coordinate.Latitude
	postOutputData.Coordinate.Longitude = post.Coordinate.Longitude
	postOutputData.CreatedAt = post.CreatedAt

	return postOutputData, nil
}
