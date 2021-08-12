package main

import (
	"Blog/config"
	"Blog/db/mysql"
	"Blog/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Init()
	config.GetMySQLConfig()
	mysql.Init()
	router.CollectRoutes(r)
	r.Run(":6243")
}
