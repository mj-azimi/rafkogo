package user

import (
	"rafkogo/utils/view"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Welcome(c *gin.Context) {
	message := h.service.GetWelcomeMessage()

	view.Render(c.Writer, 
	[]string{
		"internal/user/view/index.blade",
		"internal/user/view/header.blade",
		"internal/user/view/footer.blade",
	}, map[string]interface{}{
		"Title":   "صفحه اصلی",
		"Message": message,
		"LoggedIn": false,
	})

}
