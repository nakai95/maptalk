package port

//go:generate mockgen -source=$GOFILE -destination=../../../../mock/mock_$GOFILE -package=mock -self_package=maptalk/mock

// Controller
type UserInput interface {
	GetUserByID(id string) (UserOutputData, error)
}

// Presenter
type UserOutputData struct {
	ID   string
	Name string
}

type UserOutput interface {
	PresentUser(user UserData) (UserOutputData, error)
}

// Repository
type UserData struct {
	ID   string
	Name string
}

type UserDataAccess interface {
	FindByID(id string) (UserData, error)
}
