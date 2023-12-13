package v1

import (
	"github.com/GearFramework/emarket/internal/app"
	"github.com/GearFramework/emarket/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseCart struct {
	Cart *entities.Cart `json:"cart"`
}

func Cart(ctx *gin.Context, app app.Market) {
	ctx.JSON(http.StatusOK, ResponseCart{Cart: &entities.Cart{}})
}
