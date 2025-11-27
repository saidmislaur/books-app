package author

import (
	"strconv"

	response "books-api/internal/pkg"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOne(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid author ID", err)
	}

	auth, err := h.Service.GetOne(id)
	if err != nil {
		response.Internal(c, "Failed to fetch author", err)
	}

	if auth == nil {
		response.NotFound(c, "Author not found", err)
	}

	response.SuccessResponse(c, auth)
}
