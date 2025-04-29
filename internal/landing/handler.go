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
	title, err := h.service.GetFirstLanding()
	if err != nil {
		title = new(string)
		*title = "رکوردی یافت نشد"
	}

	view.Render(c.Writer,
		[]string{
			"internal/landing/view/index.blade",
			"internal/landing/view/header.blade",
			"internal/landing/view/footer.blade",
		}, map[string]interface{}{
			"Title":    *title,
			"Message":  message,
			"LoggedIn": false,
		})
}
