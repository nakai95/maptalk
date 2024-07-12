package usecase

import (
	"maptalk/internal/domain/entity"
	"maptalk/internal/domain/usecase/port"
)

// UseCase
type userUseCase struct {
	outputPort port.UserOutput
	dataAccess port.UserDataAccess
}

func NewUserUseCase(outputPort port.UserOutput, dataAccess port.UserDataAccess) port.UserInput {
	return &userUseCase{
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
