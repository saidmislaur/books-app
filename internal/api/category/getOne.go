package categoryapi

import (
	"strconv"

	"github.com/gin-gonic/gin"

	response "books-api/internal/pkg"
)

func (h *Handler) GetOne(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid category ID", err)
	}

	category, err := h.Service.GetOne(id)
	if err != nil {
		response.Internal(c, "Failed to fetch category: ", err)
		return
	}

	if category == nil {
		response.NotFound(c, "Category not found", err)
		return
	}
	response.SuccessResponse(c, category)
}
