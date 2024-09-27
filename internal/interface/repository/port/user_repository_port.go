package repositoryPort

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

type UserInsertData struct {
	Name   string
	Avatar string
}

type UserData struct {
	ID     string
	Name   string
	Avatar string
}
