package usecase

import (
	"maptalk/internal/domain/entity"
	"maptalk/internal/domain/usecase/port"
    "math/rand"
    "strconv"
    "context"
)

// UseCase
type userUseCase struct {
	outputPort port.UserOutput
	dataAccess port.UserDataAccess
	outputPort port.UserOutput
	dataAccess port.UserDataAccess
}

func NewUserUseCase(outputPort port.UserOutput, dataAccess port.UserDataAccess) port.UserInput {
func NewUserUseCase(outputPort port.UserOutput, dataAccess port.UserDataAccess) port.UserInput {
	return &userUseCase{
		outputPort: outputPort,
		dataAccess: dataAccess,
	}
		outputPort: outputPort,
		dataAccess: dataAccess,
	}
}

func (u *userUseCase) GetUserByID(id string) (port.UserOutputData, error) {
	userData, err := u.dataAccess.FindByID(id)
	if err != nil {
		return port.UserOutputData{}, err
	}
	// Convert user data to domain entity
	userEntity := entity.User(userData)
	// Convert domain entity to user data
	userData = port.UserData(userEntity)

	return u.outputPort.PresentUser(userData)
}

func (u *userUseCase) Save(name string, ctx context.Context) (port.UserOutputData, error) {
    user := port.UserData{
        ID: strconv.Itoa(rand.Intn(100)),
        Name: name,
    }
    n, err := u.dataAccess.Save(user, ctx)
    if err != nil {
        return port.UserOutputData{}, err
    }
    userOutputData, err := u.outputPort.PresentUser(n)
    if err != nil {
        return port.UserOutputData{}, err
    }
    return userOutputData, nil
}
