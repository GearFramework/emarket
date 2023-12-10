package backend

import (
	"github.com/GearFramework/emarket/internal/backend/handlers"
	"github.com/gin-gonic/gin"
)

func (app *Backend) initRoutes() {
	app.Server.Router.GET("/ping", func(ctx *gin.Context) {
		handlers.Ping(ctx)
	})
	app.Server.Router.NoRoute(handlers.NotFound)
}
