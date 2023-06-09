package middleware

import (
	"LiuMa-backend-go/internal/api/e"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			e.ErrorMsg(c, e.StatusUnauthorized, e.MsgFlags(e.ERROR_AUTH))
			return
		}
		c.Next()
	}
}
