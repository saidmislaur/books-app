package publisher

import (
	service "books-api/internal/service/publisher"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{Service: s}
}

func Routes(r *gin.Engine, h *Handler) {

	authors := r.Group("/authors")
	{
		authors.GET("/", h.GetAll)
		authors.GET("/:id", h.GetOne)
		authors.POST("/", h.Create)
		authors.PUT("/:id", h.Update)
		authors.DELETE("/:id", h.Delete)
	}
}
