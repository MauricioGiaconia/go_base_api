package router

import "github.com/gin-gonic/gin"

// RouteRegistrar es implementado por cada archivo de routes/ para declarar
// los endpoints de su dominio en los grupos correspondientes.
//
// public    → rutas sin autenticación (ej: POST /auth/login)
// protected → rutas que requieren JWT válido (middleware ya aplicado por el router)
//
// Ejemplo:
//
//	func (r *InvoiceRoutes) RegisterRoutes(public, protected *gin.RouterGroup) {
//		protected.GET("/invoices", r.handler.List)         // requiere JWT
//		protected.POST("/invoices", r.handler.Create)      // requiere JWT
//	}
type RouteRegistrar interface {
	RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup)
}
