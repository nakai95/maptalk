package repositoryPort

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

type PostInsertData struct {
	UserID     string
	UserName   string
	UserAvatar string
	Message    string
	Latitude   float64
	Longitude  float64
}

type PostData struct {
	ID         string
	UserID     string
	UserName   string
	UserAvatar string
	Message    string
	Latitude   float64
	Longitude  float64
	CreatedAt  int64
}
