package main

import (
	"github.com/YerzatCode/auth-service/internal/config"
	"github.com/YerzatCode/auth-service/internal/controller"
	"github.com/YerzatCode/auth-service/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// * Config init
	config.ConfigPathInit()
	config.InitConfig()

	// * Database and routes init
	database.InitDB(config.Cfg.StoragePath)
	r := gin.Default()
	controller.InitRoutes(r, database.DB)

	// * Run server
	r.Run(config.Cfg.Port)
}
