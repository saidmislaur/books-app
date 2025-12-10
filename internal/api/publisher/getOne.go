package publisher

import (
	response "books-api/internal/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOne(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid id publisher", err)
		return
	}

	publisher, err := h.Service.GetOne(id)
	if err != nil {
		response.Internal(c, "Failed to fetch one publisher", err)
		return
	}

	if publisher == nil {
		response.NotFound(c, "publisher not found", err)
		return
	}

	response.SuccessResponse(c, publisher)
}
