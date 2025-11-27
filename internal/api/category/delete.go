package categoryapi

import (
	"strconv"

	response "books-api/internal/pkg"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid category ID", err)
		return
	}

	err = h.Service.Delete(id)
	if err != nil {
		response.Internal(c, "Failed to delete category", err)
		return
	}

	response.SuccessResponse(c, gin.H{
		"message": "Category deleted successfully",
		"id":      id,
	})
}
