# Router

Contiene la **configuración de rutas** de la API y la interfaz de registro.

## Responsabilidades

- Definir la interfaz `RouteRegistrar` que los archivos de `routes/` implementan.
- Crear los grupos de rutas base (`public` y `protected` con JWT).
- Delegar el registro de rutas específicas a cada `RouteRegistrar`.
- Mantener las rutas genéricas de infraestructura (`/ping`, `NoRoute`).

## Reglas

- Las rutas específicas de cada dominio se definen en `routes/<dominio>_routes.go`.
- El router es el orquestador: crea los grupos y delega. **Nunca se modifica al agregar un nuevo dominio.**
- Nomenclatura: función `SetupRoutes(r *gin.Engine, registrars []RouteRegistrar)`.
- Archivos: `router.go` (setup), `types.go` (interfaz).

## Interfaz RouteRegistrar

```go
type RouteRegistrar interface {
    RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup)
}
```

- `public` → grupo sin autenticación (ej: `/auth/login`, `/auth/register`)
- `protected` → grupo con `middleware.RequireAccessToken` ya aplicado (ej: `/invoices`, `/users/me`)

Los archivos de `routes/` satisfacen esta interfaz de forma implícita (structural typing de Go), sin necesidad de importar el paquete `router`.
