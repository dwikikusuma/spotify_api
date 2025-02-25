package memberships

import (
	"github.com/gin-gonic/gin"
	"spotify_api/internal/model/memberhsips"
)

type Service interface {
	SignUp(request memberhsips.SignUpRequest) error
}

type Handler struct {
	*gin.Engine
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/users")
	route.POST("/sign-up", h.SignUp)
}
