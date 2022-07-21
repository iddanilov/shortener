package main

import (
	"github.com/gin-gonic/gin"

	"github.com/shortener/internal/app/handlers"
	"github.com/shortener/internal/app/storage"
)

func main() {
	r := gin.New()

	storage := storage.NewStorage()
	rg := handlers.NewRouterGroup(&r.RouterGroup, &storage)
	rg.Routes()

	r.Run(":8080")
}
