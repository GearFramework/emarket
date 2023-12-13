package backend

import (
	"github.com/GearFramework/emarket/internal/backend/handlers"
	v1 "github.com/GearFramework/emarket/internal/backend/handlers/v1"
	"github.com/gin-gonic/gin"
)

func (app *Backend) initRoutes() {
	app.Server.Router.GET("/v1/cart", func(ctx *gin.Context) {
		v1.Cart(ctx, app)
	})
	app.Server.Router.GET("/ping", func(ctx *gin.Context) { handlers.Ping(ctx) })
	app.Server.Router.NoRoute(handlers.NotFound)
}
