package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, ErrorResponse{Error: message})
}

func BadRequest(c *gin.Context, msg string, err error) {
	Error(c, http.StatusBadRequest, msg+": "+err.Error())
}

func NotFound(c *gin.Context, msg string, err error) {
	Error(c, http.StatusNotFound, msg+": "+err.Error())
}

func Internal(c *gin.Context, action string, err error) {
	Error(c, http.StatusInternalServerError, action+": "+err.Error())
}

//успешная обработка

func SuccessResponse(c *gin.Context, data any) {
	if data == nil {
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusOK, data)
}
