package publisher

import (
	response "books-api/internal/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid id publisher", err)
		return
	}

	err = h.Service.Delete(id)
	if err != nil {
		response.Internal(c, "Invalit delete publisher", err)
	}

	response.SuccessResponse(c, gin.H{
		"message": "Издательство удалено",
	})
}
