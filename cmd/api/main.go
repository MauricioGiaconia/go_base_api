package main

import (
	"github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/router"
	"github.com/MauricioGiaconia/go_base_api/internal/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	c := bootstrap.NewContainer()

	r := gin.Default()
	router.SetupRoutes(r, c.Registrars)
	r.Run(":8080")
}
