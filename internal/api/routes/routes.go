package routes

import (
	"LiuMa-backend-go/internal/api/handlers"
	"LiuMa-backend-go/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SystemRouter(r *gin.Engine) {
	r.POST("/autotest/register", handlers.Register)
	r.POST("/autotest/login", handlers.Login)
	//r.POST("/login", system.Login)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())

}
