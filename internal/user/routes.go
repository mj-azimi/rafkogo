package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	service := NewUserService()
	handler := NewUserHandler(service)

	r.GET("/", handler.Welcome)
}
