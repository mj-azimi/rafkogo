package bootstrap

import (
	"rafkogo/config"
	"rafkogo/routes"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	config.LoadEnv()
	r := gin.Default()
	r.Static("/public", "./public")
	routes.RegisterRoutes(r)

	return r
}
