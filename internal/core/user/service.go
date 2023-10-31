package user

import (
	"errors"
	"hello-api-go/internal/entity"
	"hello-api-go/internal/infra"

	"go.uber.org/zap"
)

type UserService struct {
	userRepository infra.UserRepository
	logger zap.Logger
}

func (u UserService) Create(createUserRequest CreateUserRequest) (entity.User, error) {
	u.logger.Info("Creating a user from request", zap.Any("request", createUserRequest));
	
	if !createUserRequest.Validate() {
		return entity.User{}, errors.New("Error while validating request body");
	}

	user, err := u.userRepository.Add(createUserRequest.ToUserEntity());
	if err != nil {
		return entity.User{}, err;
	}

	return user, nil;
}

func MakeUserService(userRepository infra.UserRepository, logger zap.Logger) UserService {
	return UserService{userRepository: userRepository, logger: logger};
}