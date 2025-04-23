package routes

import (
	"github.com/gin-gonic/gin"

	"rafkogo/internal/user"
	"rafkogo/internal/landing"
)

func RegisterRoutes(r *gin.Engine) {

	
	landingRoute := r.Group("/")
	landing.RegisterRoutes(landingRoute.Group("/"))

	api := r.Group("/api")
	user.RegisterRoutes(api.Group("/users"))
}
