package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Welcome(c *gin.Context) {
	message := h.service.GetWelcomeMessage()
	c.JSON(http.StatusOK, gin.H{"message": message})
}
