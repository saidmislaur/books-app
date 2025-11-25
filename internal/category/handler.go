package category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cm *CategoryManager) GetCategoriesHandler(c *gin.Context) {
	categories, err := cm.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch authors: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (cm *CategoryManager) GetCategoryHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category ID",
		})
	}

	category, err := cm.GetCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch category: " + err.Error(),
		})
		return
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cm *CategoryManager) CreateCategoryHandler(c *gin.Context) {
	var req Category

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request body: " + err.Error(),
		})
		return
	}

	category, err := cm.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create new category" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (cm *CategoryManager) UpdateCategoryHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	var req Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalud request body" + err.Error(),
		})
		return
	}

	_, err = cm.UpdateCategory(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "категория обновлена",
		"id":      id,
	})
}

func (cm *CategoryManager) DeleteCategoryHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	err = cm.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
		"id":      id,
	})
}
