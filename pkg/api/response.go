package api

import (
	"hello-api-go/pkg/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendData(ctx *gin.Context, result data.Result) {
	switch result.Error {
	case data.EmptyResult:
		ctx.Status(http.StatusNoContent);
	case data.ConflictData:
		ctx.Status(http.StatusConflict);
	case data.GenericError:
		if result.Result != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"data": result.Result});
			return
		}
		ctx.Status(http.StatusUnprocessableEntity);
	case data.ValidationError:
		if result.Result != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data": result.Result});
			return
		}
		ctx.Status(http.StatusBadRequest);
	default:
		if result.Result != nil {
			ctx.JSON(http.StatusOK, gin.H{"data": result.Result});
			return
		}
		ctx.Status(http.StatusOK);
	}
}

