package user

import (
	"hello-api-go/internal/configuration"
	"hello-api-go/internal/infra"
	"hello-api-go/pkg/api"
	"hello-api-go/pkg/data"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserController struct {
	userService UserService
	logger zap.Logger
}

func (controller UserController) CreateUserHandler(context *gin.Context) {
	var requestBody CreateUserRequest;
	context.BindJSON(&requestBody);

	validate := validator.New();
	err := validate.Struct(requestBody);

	if err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessages := api.MakeMessageFromFieldError(errors);
		api.SendData(context, data.Result{Result: gin.H{
			"error": true,
			"validation": errorMessages,
		}, Error: data.ValidationError});
		return
	}

	result := controller.userService.Create(requestBody);

	api.SendData(context, result);
}

func (controller UserController) FindUserHandler(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("id"));
	if err != nil {
		api.SendData(context, data.Result{Result: gin.H{
			"error": true,
			"message": "Invalid user id",
		}, Error: data.ValidationError});
		return;
	}
	result := controller.userService.Find(userId);
	api.SendData(context, result);
}

func (controller UserController) RegisterEndpoints(router *gin.Engine) {
	router.POST("/users", controller.CreateUserHandler);
	router.GET("/users/:id", controller.FindUserHandler);
}

func MakeUserController(envObject configuration.Env) UserController {
	userRepository := infra.MakeUserRepository(envObject.Db);
	userService := MakeUserService(userRepository, envObject.Logger);
	return UserController{userService: userService, logger: envObject.Logger};
}

