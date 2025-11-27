package categoryapi

import (
	"strconv"

	"books-api/internal/models"

	response "books-api/internal/pkg"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid category ID", err)
		return
	}

	var req models.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Internal(c, "Invalud request body", err)
		return
	}

	_, err = h.Service.Update(id, req)
	if err != nil {
		response.Internal(c, "Failed to update category", err)
		return
	}

	response.SuccessResponse(c, gin.H{
		"message": "категория обновлена",
		"id":      id,
	})

}
