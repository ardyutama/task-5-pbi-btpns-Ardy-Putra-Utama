package main

import (
	"rakamin-golang/database"
	"rakamin-golang/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()
	router.SetupRoutes(r)
	r.Run(":8080")
}
