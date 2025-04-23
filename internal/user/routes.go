package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.Use(LoggerMiddleware())

	service := NewUserService()
	handler := NewUserHandler(service)
	r.GET("/login", handler.Welcome)
}
