package categoryapi

import (
	service "books-api/internal/service/category"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{Service: s}
}

//routes

func RegisterRoutes(r *gin.Engine, h *Handler) {
	group := r.Group("/categories")

	group.GET("/", h.GetAll)
	group.GET("/:id", h.GetOne)
	group.POST("/", h.Create)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}
