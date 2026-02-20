# Router

Contiene la **configuración centralizada de rutas** de la API.

## Responsabilidades

- Registrar todos los endpoints de la API.
- Aplicar middlewares a los grupos de rutas correspondientes.
- Conectar los handlers con sus rutas HTTP.

## Reglas

- Es el único lugar donde se definen rutas.
- Recibe los handlers como dependencias inyectadas desde `main.go`.
- Nomenclatura: función `SetupRoutes(r *gin.Engine, ...)`.
- Archivo: `router.go`.
