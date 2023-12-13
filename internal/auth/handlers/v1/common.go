package v1

import (
	"github.com/GearFramework/emarket/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, models.ResponseInvalidLogin{
		models.Response{Status: models.Ok},
		uint32(400),
		message,
	})
}
