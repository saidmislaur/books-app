package author

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db *sql.DB) {
	manager := NewManager(db)

	authors := r.Group("/authors")
	{
		authors.GET("/", manager.GetAllAuthors)
		authors.GET("/:id", manager.GetOneAuthor)
		authors.POST("/", manager.Create)
		authors.PUT("/:id", manager.Update)
		authors.DELETE("/:id", manager.Delete)
	}
}
