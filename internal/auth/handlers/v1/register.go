package v1

import (
	"encoding/json"
	"github.com/GearFramework/emarket/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context, a models.Identifier) {
	defer ctx.Request.Body.Close()
	dec := json.NewDecoder(ctx.Request.Body)
	var req models.RequestRegister
	if err := dec.Decode(&req); err != nil {
		LoginBadRequest(ctx, "Login: bad request")
		return
	}
	a.Register(ctx, &req)
	ctx.JSON(http.StatusOK, struct {
	}{})
}
