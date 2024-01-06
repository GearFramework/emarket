package middleware

import (
	"fmt"
	"github.com/GearFramework/emarket/internal/models"
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	CookieParamName = "Authorization"
	TokenExpired    = time.Hour * 24
)

var logger *alog.Alog

func Auth(a models.Identifier) gin.HandlerFunc {
	logger = alog.NewLogger("IdentityMiddleware")
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie(CookieParamName)
		a.IdentityByCookie(cookie)
		if err != nil || cookie == nil || cookie.Value == "" {
		}
		fmt.Println(err)
		//if token == "" {
		//	logger.Warnf("%s request: %s; status: %d; size: %d; Empty authorization token",
		//		ctx.Request.Method,
		//		ctx.Request.RequestURI,
		//		ctx.Writer.Status(),
		//		ctx.Writer.Size(),
		//	)
		//	//UnauthorizedResponse.Send(ctx)
		//	return
		//}
		ctx.Next()
	}
}

//
//func authFromToken(ctx *gin.Context, token string) (string, error) {
//	userID, err := s.api.Auth(token)
//	if err != nil && err == auth.ErrInvalidAuthorization {
//		ctx.AbortWithStatus(http.StatusUnauthorized)
//		return 0, err
//	}
//	if err != nil && err == auth.ErrNeedAuthorization {
//		return s.AuthNewUser(ctx)
//	}
//	s.setAuthCookie(ctx, token)
//	return userID, err
//}
//
//func setAuthCookie(ctx *gin.Context, token string) {
//	ctx.SetCookie(CookieParamName,
//		token,
//		int(TokenExpired.Seconds()),
//		"/",
//		"localhost",
//		true,
//		true,
//	)
//}
