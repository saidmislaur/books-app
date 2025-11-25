package category

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db *sql.DB) {
	manager := NewManager(db)

	categories := r.Group("/categories")
	{
		categories.GET("/", manager.GetCategoriesHandler)
		categories.GET("/:id", manager.GetCategoryHandler)
		categories.POST("/", manager.CreateCategoryHandler)
		categories.PUT("/:id", manager.UpdateCategoryHandler)
		categories.DELETE("/:id", manager.DeleteCategoryHandler)
	}
}
