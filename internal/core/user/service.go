package user

import (
	"hello-api-go/internal/infra"
	"hello-api-go/pkg/data"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type UserService struct {
	userRepository infra.UserRepository
	logger zap.Logger
}

func (u UserService) Create(createUserRequest CreateUserRequest) data.Result {
	u.logger.Info("Creating a user from request", zap.Any("request", createUserRequest));
	user, err := u.userRepository.Add(createUserRequest.ToUserEntity());
	if err != nil {
		u.logger.Error("Error while creating user", zap.Error(err));
		return data.Result{Error: data.GenericError, Result: nil};
	}
	return data.Result{Result: user};
}

func (u UserService) Find(userId int) data.Result {
	user, err := u.userRepository.FindById(userId);
	if err != nil {
		u.logger.Error("Error while querying user", zap.Error(err));
		if err == pgx.ErrNoRows {
			return data.Result{Error: data.EmptyResult, Result: nil};
		}
		return data.Result{Error: data.GenericError, Result: nil};
	}
	result := fromUserEntity(user);
	return data.Result{Result: result};
}

func MakeUserService(userRepository infra.UserRepository, logger zap.Logger) UserService {
	return UserService{userRepository: userRepository, logger: logger};
}