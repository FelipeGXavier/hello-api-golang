package main

import (
	"hello-api-go/internal/configuration"
	"hello-api-go/internal/core/user"

	"github.com/gin-gonic/gin"
)



func main() {
	config, err := configuration.LoadConfig();
	if err != nil {
		panic(err);
	}

	env := *configuration.LoadEnvObject(config);

	userControler := user.MakeUserController(env);
	
	router := gin.Default();
	router.POST("/users", userControler.CreateUserHandler);
	router.Run();
}