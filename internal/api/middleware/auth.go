package middleware

import (
	"LiuMa-backend-go/internal/api/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": e.MsgFlags(e.ERROR_AUTH)})
			c.Abort()
			return
		}

		c.Next()
	}
}
