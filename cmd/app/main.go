package main

import (
	"books-api/internal/author"
	"books-api/internal/category"
	"books-api/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.Connect()

	author.Routes(r, db)
	category.Routes(r, db)

	r.Run(":5000")
}
