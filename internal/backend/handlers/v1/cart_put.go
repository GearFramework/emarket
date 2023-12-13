package v1

import (
	"github.com/GearFramework/emarket/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutToCart(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ResponseCart{Cart: &entities.Cart{}})
}
