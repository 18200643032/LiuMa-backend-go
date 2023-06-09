package main

import (
	"LiuMa-backend-go/config"
	"LiuMa-backend-go/internal/api/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	gin.SetMode(config.Config.BaseModel)
	r := routes.SystemRouter()
	if err := http.ListenAndServe(strings.Join([]string{config.Config.BaseHost,
		":", config.Config.BasePort}, ""), r); err != nil {
		panic("Port occupancy!!!")
	}

}
