package bootstrap

import (
	"database/sql"

	"github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/router"
	"github.com/MauricioGiaconia/go_base_api/internal/adapter/outbound/mysql"
)

// Container centraliza todas las dependencias de la aplicación.
// Actúa como composition root: crea y conecta todas las piezas concretas.
// Es el único lugar que conoce las implementaciones de los adapters.
type Container struct {
	DB         *sql.DB
	Registrars []router.RouteRegistrar
}

// NewContainer inicializa la infraestructura, servicios y handlers,
// retornando un contenedor con todas las dependencias listas para usar.
func NewContainer() *Container {
	// --- Infraestructura compartida ---
	db := mysql.NewConnection()

	// --- Módulos de dominio ---
	return &Container{
		DB:         db,
		Registrars: []router.RouteRegistrar{},
	}
}
