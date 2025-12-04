package main

import (
	handlerAuth "books-api/internal/api/author"
	handler "books-api/internal/api/category"
	"books-api/internal/database"
	repositoryAuth "books-api/internal/repository/author"
	repository "books-api/internal/repository/category"
	serviceAuth "books-api/internal/service/author"
	service "books-api/internal/service/category"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.Connect()

	repoAuth := repositoryAuth.NewManager(db)
	serviceAuth := serviceAuth.New(repoAuth)
	authorHandler := handlerAuth.New(serviceAuth)

	handlerAuth.Routes(r, authorHandler)

	repo := repository.New(db)
	service := service.New(repo)
	categoryHandler := handler.New(service)

	handler.RegisterRoutes(r, categoryHandler)

	r.Run(":5000")
}
