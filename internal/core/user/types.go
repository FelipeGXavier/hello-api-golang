package user

import "hello-api-go/internal/entity"

type UserResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}

func fromUserEntity(user entity.User) UserResponse {
	return UserResponse{
		Id: user.Id,
		Name: user.Name,
		Surname: user.Surname,
		Email: user.Email,
	}
}

type CreateUserRequest struct {
	Name string `validate:"required,min=3,max=30"`
	Surname string `validate:"required,min=3,max=50"`
	Email string `validate:"required,min=5,max=60"`
}

func (request CreateUserRequest) ToUserEntity() entity.User {
	return entity.User{
		Name: request.Name,
		Surname: request.Surname,
		Email: request.Email,
	};
}