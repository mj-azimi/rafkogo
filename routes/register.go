package routes

import (
	"github.com/gin-gonic/gin"

	"rafkogo/internal/user"     
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to your Go project!"})
	})

	api := r.Group("/api")
	user.RegisterRoutes(api.Group("/users"))
}
