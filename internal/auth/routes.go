package auth

import (
	v1 "github.com/GearFramework/emarket/internal/auth/handlers/v1"
	"github.com/GearFramework/emarket/internal/backend/handlers"
	"github.com/GearFramework/emarket/internal/pkg/auth"
	"github.com/gin-gonic/gin"
)

func (app *ServiceAuth) initRoutes() {
	a := auth.Auth{
		SecretKey: app.Config.AuthKey,
	}
	app.Server.Router.GET("/v1/login", func(ctx *gin.Context) {
		v1.Login(ctx, a)
	})
	app.Server.Router.GET("/ping", func(ctx *gin.Context) { handlers.Ping(ctx) })
	app.Server.Router.NoRoute(handlers.NotFound)
}
