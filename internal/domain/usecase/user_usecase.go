package usecase

import (
	"maptalk/internal/domain/entity"
)

type UserInputPort interface {
	GetUserByID(id string) (*UserOutputData, error)
}

type UserOutputData struct {
    ID   string
    Name string
}

type UserOutputPort interface {
    PresentUser(user *UserData) (*UserOutputData, error)
}

type UserData struct {
    ID   string
    Name string
}

type UserDataAccess interface {
    FindByID(id string) (*UserData, error)
}

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