package author

import (
	"books-api/internal/models"
	response "books-api/internal/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid author ID", err)
		return
	}

	var req models.Author
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Internal(c, "Invalid request body", err)
		return
	}

	_, err = h.Service.Update(id, req)
	if err != nil {
		response.Internal(c, "invalid to update author", err)
		return
	}

	response.SuccessResponse(c, gin.H{
		"message": "автор обнавлен",
		"id":      id,
	})
}
