package author

import (
	"books-api/internal/models"
	response "books-api/internal/pkg"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var req models.Author

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Failed request body", err)
		return
	}

	author, err := h.Service.Create(req)
	if err != nil {
		response.Internal(c, "Failed to cretae new author", err)
		return
	}

	response.SuccessResponse(c, author)
}
