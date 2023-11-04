package api

import (
	"fmt"
	"hello-api-go/pkg/functions"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TranslateFieldError(field validator.FieldError) string {
	switch field.Tag() {
	case "required":
		return fmt.Sprintf("The field %s is required", field.Field());
	case "min":
		return fmt.Sprintf("The field %s should be greater than %s characters", field.Field(), field.Param());
	case "max":
		return fmt.Sprintf("The field %s should be less than %s characters", field.Field(), field.Param());
	default:
		return fmt.Sprintf("The field %s is invalid", field.Field());;
	}
}

func MakeMessageFromFieldError(errors validator.ValidationErrors) []gin.H {
	
	r := functions.Map[validator.FieldError, gin.H](errors, func(fe validator.FieldError) gin.H {
		return gin.H{
			"field": fe.Field(),
			"error": TranslateFieldError(fe),
		};
	});

	for _, x := range r {
		fmt.Println(x);
	}

	return r;
}