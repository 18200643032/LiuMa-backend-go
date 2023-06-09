package routes

import (
	"LiuMa-backend-go/internal/api/handlers"
	"LiuMa-backend-go/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SystemRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]any{"code": "200", "msg": "ok"})
	})
	v1 := r.Group("/autotest")
	{
		v1.POST("/register", handlers.Register)
		v1.POST("/login", handlers.Login)
	}
	//r.POST("/login", system.Login)
	v2 := r.Group("/api/v1")
	v2.Use(middleware.AuthMiddleware())
	return r

}
