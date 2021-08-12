package main

import (
	"Blog/config"
	"Blog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Init()
	config.GetMySQLConfig()
	// mysql.Init()
	router.CollectRoutes(r)
	r.Run(":6243")
}
