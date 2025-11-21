package author

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (am *AuthorManager) GetAllAuthors(c *gin.Context) {
	authors, err := am.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch authors: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (am *AuthorManager) GetOneAuthor(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid author ID",
		})
	}

	author, err := am.GetAuthorById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch author: " + err.Error(),
		})
		return
	}

	if author == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Author not found",
		})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (am *AuthorManager) Create(c *gin.Context) {
	var req Author

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	author, err := am.CreateAuthor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create author: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func (am *AuthorManager) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid author ID",
		})
		return
	}

	var req Author

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	_, err = am.UpdateAuthor(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create author: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author updated successfully",
		"id":      id,
	})
}

func (am *AuthorManager) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid author ID",
		})
		return
	}

	err = am.DeleteAuthor(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete author: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author deleted successfully",
		"id":      id,
	})
}
