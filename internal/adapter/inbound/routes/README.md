# Routes

Contiene la **declaración de rutas por dominio**, donde cada archivo registra los endpoints de un módulo específico.

## Responsabilidades

- Definir un `struct` por dominio (ej: `AuthRoutes`, `InvoiceRoutes`) que agrupa el handler correspondiente.
- Implementar el método `RegisterRoutes` para montar los endpoints en los grupos `public` y/o `protected`.
- Delegar la lógica HTTP al handler declarado en el struct, sin incluir lógica de negocio.

## Reglas

- Un archivo por dominio: `<dominio>_routes.go` (ej: `auth_routes.go`, `invoice_routes.go`).
- Cada struct implementa `router.RouteRegistrar` de forma implícita via structural typing, **sin importar el paquete `router`**.
- Las rutas públicas (sin autenticación) se montan sobre el grupo `public` (ej: `/auth/login`).
- Las rutas protegidas (requieren JWT) se montan sobre el grupo `protected` (ej: `/invoices`, `/users/me`).
- **Nunca** se incluye lógica de negocio ni acceso a datos en este paquete.
- El constructor sigue la nomenclatura `New<Dominio>Routes(h *handler.<Dominio>Handler) *<Dominio>Routes`.

## Estructura de un archivo de rutas

```go
package routes

import (
    "github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/handler"
    "github.com/gin-gonic/gin"
)

// <Dominio>Routes declara los endpoints del dominio de <dominio>.
// Implementa router.RouteRegistrar de forma implícita via structural typing,
// desacoplando la declaración de rutas de la lógica HTTP del handler.
type <Dominio>Routes struct {
    handler *handler.<Dominio>Handler
}

// New<Dominio>Routes crea el registrar de rutas para el dominio de <dominio>.
func New<Dominio>Routes(h *handler.<Dominio>Handler) *<Dominio>Routes {
    return &<Dominio>Routes{handler: h}
}

// RegisterRoutes monta los endpoints de <dominio> en los grupos correspondientes.
func (r *<Dominio>Routes) RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
    // Rutas públicas
    public.Group("/<dominio>").POST("/...", r.handler.MetodoPublico)

    // Rutas protegidas
    protected.Group("/<dominio>").GET("/...", r.handler.MetodoProtegido)
}
```

## Ejemplo: `auth_routes.go`

```go
package routes

import (
    "github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/handler"
    "github.com/gin-gonic/gin"
)

type AuthRoutes struct {
    handler *handler.AuthHandler
}

func NewAuthRoutes(h *handler.AuthHandler) *AuthRoutes {
    return &AuthRoutes{handler: h}
}

func (r *AuthRoutes) RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
    auth := public.Group("/auth")
    {
        auth.POST("/login", r.handler.Login)
        auth.POST("/register", r.handler.Register)
    }
}
```

## Integración con el router

Los registrars se instancian en el bootstrap/contenedor y se pasan a `router.SetupRoutes`:

```go
registrars := []router.RouteRegistrar{
    routes.NewAuthRoutes(authHandler),
    routes.NewInvoiceRoutes(invoiceHandler),
}

router.SetupRoutes(r, registrars)
```

El router **nunca se modifica** al agregar un nuevo dominio; basta con crear el archivo `<dominio>_routes.go` e inyectar el registrar en el bootstrap.
