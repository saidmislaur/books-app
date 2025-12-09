package publisher

import (
	response "books-api/internal/pkg"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	list, err := h.Service.GetAll()
	if err != nil {
		response.Internal(c, "Failed to fetch publishers", err)
		return
	}

	response.SuccessResponse(c, list)
}
