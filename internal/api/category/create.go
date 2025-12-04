package categoryapi

import (
	"books-api/internal/models"
	response "books-api/internal/pkg"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var req models.Category

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Failed request body: ", err)
		return
	}

	category, err := h.Service.Create(req)
	if err != nil {
		response.Internal(c, "Failed to create new category", err)
		return
	}

	response.SuccessResponse(c, category)
}
