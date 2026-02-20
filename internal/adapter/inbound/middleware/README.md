# Middleware

Contiene los **middlewares HTTP** que interceptan las peticiones antes de llegar a los handlers.

## Responsabilidades

- Validar tokens JWT de acceso (`RequireAccessToken`).
- Validar tokens de contrato contra la base de datos (`ContractAuthMiddleware`).
- Establecer datos en el contexto de Gin para uso downstream.

## Reglas

- Los middlewares que necesitan acceso a datos deben depender de **output ports** (interfaces), no de implementaciones concretas.
- Nomenclatura para funciones middleware: verbos descriptivos (ej: `RequireAccessToken`).
- Nomenclatura para factory functions: `<Dominio>AuthMiddleware(repo port.SomePort) gin.HandlerFunc`.
- Archivo: `<concepto>_middleware.go` en snake_case.
