package router

import (
	"net/http"

	"github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/response"
	"github.com/gin-gonic/gin"
)

// SetupRoutes centraliza la configuración de todas las rutas de la API.
// Recibe el router de Gin y los handlers necesarios como dependencias inyectadas.
func SetupRoutes(r *gin.Engine) {

	// Endpoint ping para verificar el funcionamiento de la API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.ToApi(http.StatusOK, "Pong", false, 0, 0, 0))
	})

	// Maneja rutas no encontradas
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ToApi(http.StatusNotFound, "endpoint not found", false, 0, 0, 0))
	})
}
