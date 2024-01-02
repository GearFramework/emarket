package v1

import (
	"github.com/GearFramework/emarket/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, models.ResponseInvalidLogin{
		Response:     models.Response{Status: models.Ok},
		ErrorCode:    uint32(400),
		ErrorMessage: message,
	})
}

func RegisterBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, models.ResponseInvalidRegister{
		Response:     models.Response{Status: models.Ok},
		ErrorCode:    uint32(400),
		ErrorMessage: message,
	})
}
