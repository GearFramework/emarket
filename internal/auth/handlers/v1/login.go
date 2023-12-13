package v1

import (
	"encoding/json"
	"github.com/GearFramework/emarket/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context, a models.Identifier) {
	defer ctx.Request.Body.Close()
	dec := json.NewDecoder(ctx.Request.Body)
	var req models.RequestLogin
	if err := dec.Decode(&req); err != nil {
		LoginBadRequest(ctx, "Login: bad request")
		return
	}
	a.Login(ctx, &req)
	ctx.JSON(http.StatusOK, struct {
	}{})
}
