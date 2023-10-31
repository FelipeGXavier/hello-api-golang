package user

import (
	"hello-api-go/internal/configuration"
	"hello-api-go/internal/entity"
	"hello-api-go/internal/infra"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name string
	Surname string
	Email string
}

func (request CreateUserRequest) Validate() bool {
	return true;
}

func (request CreateUserRequest) ToUserEntity() entity.User {
	return entity.User{
		Name: request.Name,
		Surname: request.Surname,
		Email: request.Email,
	};
}

type UserController struct {
	userService UserService
}

func (controller UserController) CreateUserHandler(context *gin.Context) {
	var requestBody CreateUserRequest;
	context.BindJSON(&requestBody);
	user, err := controller.userService.Create(requestBody);
	if err != nil {
		context.AbortWithStatusJSON(400, gin.H{
			"error": true,
		});
	}
	context.JSON(http.StatusCreated, gin.H{
		"data": user,
	});
}

func MakeUserController(envObject configuration.Env) UserController {
	userRepository := infra.MakeUserRepository(envObject.Db);
	userService := MakeUserService(userRepository, envObject.Logger);
	return UserController{userService: userService};
}

