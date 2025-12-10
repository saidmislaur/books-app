package publisher

import (
	"books-api/internal/models"
	response "books-api/internal/pkg"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var req models.Publisher

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Failed request body", err)
		return
	}

	publisher, err := h.Service.Create(req)
	if err != nil {
		response.Internal(c, "Failed to create new publisher", err)
		return
	}

	response.SuccessResponse(c, publisher)
}
