package main

import (
	handlerAuth "books-api/internal/api/author"
	handler "books-api/internal/api/category"
	handlerPublisher "books-api/internal/api/publisher"
	"books-api/internal/database"
	repositoryAuth "books-api/internal/repository/author"
	repository "books-api/internal/repository/category"
	repositoryPublisher "books-api/internal/repository/publisher"
	serviceAuth "books-api/internal/service/author"
	service "books-api/internal/service/category"
	servicePublisher "books-api/internal/service/publisher"

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

	repoPublish := repositoryPublisher.NewManager(db)
	servicePublisher := servicePublisher.New(repoPublish)
	publisherHandler := handlerPublisher.New(servicePublisher)

	handlerPublisher.Routes(r, publisherHandler)

	r.Run(":5000")
}
