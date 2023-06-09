package main

import (
	"LiuMa-backend-go/internal/api/routes"
	"LiuMa-backend-go/internal/database"
)

func main() {
	database.InitMysql()
	r := routes.InitRouter()
	r.Run(":9999")

}
