package middleware

import (
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenCheck(c *gin.Context) {
	appG := common.Gin{C: c}
	token := c.Request.Header.Get("token")
	addr, ok := service.TokenCheck(token)
	if !ok {
		appG.Response(http.StatusUnauthorized, common.AUTH_TOKEN_ERROR, nil)
		c.Abort()
		return
	}
	c.Set("address", addr)
	c.Next()
}
