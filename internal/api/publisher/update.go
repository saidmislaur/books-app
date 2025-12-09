package publisher

import (
	"books-api/internal/models"
	response "books-api/internal/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	var req models.Publisher
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid id publiher", err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Internal(c, "Invalid request body", err)
		return
	}

	_, err = h.Service.Update(id, req)
	if err != nil {
		response.Internal(c, "Failed to update publisher", err)
		return
	}

	response.SuccessResponse(c, gin.H{
		"message": "Издательство обнавлено",
		"req":     req,
	})
}
