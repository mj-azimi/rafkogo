package landing

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.Use(LoggerMiddleware())

	service := NewLandingService()
	handler := NewlandingHandler(service)
	r.GET("/", handler.Welcome)
}
