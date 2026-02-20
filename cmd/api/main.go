package main

import (
	"github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// c := bootstrap.NewContainer()

	r := gin.Default()
	router.SetupRoutes(r)
	r.Run(":8080")
}
