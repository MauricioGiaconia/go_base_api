package router

import (
	"net/http"

	"github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/response"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas genéricas de la API y delega el registro
// de rutas específicas a cada RouteRegistrar.
// Recibe el router de Gin y un slice de registrars que auto-registran sus rutas.
func SetupRoutes(r *gin.Engine, registrars []RouteRegistrar) {

	// Endpoint ping para verificar el funcionamiento de la API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.ToApi(http.StatusOK, "Pong", false, 0, 0, 0))
	})

	// Maneja rutas no encontradas
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ToApi(http.StatusNotFound, "endpoint not found", false, 0, 0, 0))
	})

	// Grupos base: público y protegido
	public := r.Group("/")
	protected := r.Group("/")

	// Cada registrar define sus propias rutas dentro de los grupos que necesite
	for _, reg := range registrars {
		reg.RegisterRoutes(public, protected)
	}
}
