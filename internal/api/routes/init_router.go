package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	SystemRouter(r)
	return r
}
