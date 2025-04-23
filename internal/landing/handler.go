package landing

import (
	"rafkogo/utils/view"

	"github.com/gin-gonic/gin"
)

type landingHandler struct {
	service LandingService
}

func NewlandingHandler(service LandingService) *landingHandler {
	return &landingHandler{service: service}
}

func (h *landingHandler) Welcome(c *gin.Context) {
	message := h.service.GetWelcomeMessage()

	view.Render(c.Writer,
		[]string{
			"internal/landing/view/index.blade",
			"internal/landing/view/header.blade",
			"internal/landing/view/footer.blade",
		}, map[string]interface{}{
			"Title":    "صفحه اصلی",
			"Message":  message,
			"LoggedIn": false,
		})

}
