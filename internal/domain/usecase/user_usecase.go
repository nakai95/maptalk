package usecase

import (
	"maptalk/internal/domain/entity"
    "math/rand"
    "strconv"
    "context"
)

// Controller
type UserInputPort interface {
	GetUserByID(id string) (*UserOutputData, error)
    Save(name string, ctx context.Context) (*UserOutputData, error)
}

// Presenter
type UserOutputData struct {
    ID   string
    Name string
}

type UserOutputPort interface {
    PresentUser(user *UserData) (*UserOutputData, error)
}

// Repository
type UserData struct {
    ID   string
    Name string
}

type UserDataAccess interface {
    FindByID(id string) (*UserData, error)
    Save(user UserData, ctx context.Context) (*UserData, error)
}

// UseCase
type userUseCase struct {
	outputPort UserOutputPort
    dataAccess UserDataAccess
}

func NewUserUseCase(outputPort UserOutputPort, dataAccess UserDataAccess) UserInputPort {
	return &userUseCase{
        outputPort: outputPort,
        dataAccess: dataAccess,
    }
}
func (u *userUseCase) GetUserByID(id string) (*UserOutputData, error) {
	userData, err := u.dataAccess.FindByID(id)
    if err != nil {
        return nil, err
    }
    // Convert user data to domain entity
    userEntity := &entity.User{
        ID:   userData.ID,
        Name: userData.Name,
    }
    // Convert domain entity to user data
    userData = &UserData{
        ID:   userEntity.ID,
        Name: userEntity.Name,
    }

   return u.outputPort.PresentUser(userData)
}

func (u *userUseCase) Save(name string, ctx context.Context) (*UserOutputData, error) {
    user := &UserData{
        ID: strconv.Itoa(rand.Intn(100)),
        Name: name,
    }
    n, err := u.dataAccess.Save(*user, ctx)
    if err != nil {
        return nil, err
    }
    userOutputData, err := u.outputPort.PresentUser(n)
    if err != nil {
        return nil, err
    }
    return userOutputData, nil
}