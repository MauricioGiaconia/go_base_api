# Handler

Contiene los **controladores HTTP** (adaptadores inbound) de la API.

## Responsabilidades

- Recibir y validar las peticiones HTTP.
- Traducir los datos de la request a entidades de dominio.
- Invocar los input ports (interfaces de servicio) para ejecutar la lógica de negocio.
- Construir y retornar la respuesta HTTP.

## Reglas

- Dependen de los **input ports** definidos en `domain/port/`, NO de implementaciones concretas.
- Las dependencias se inyectan en el constructor.
- Nomenclatura: `<Dominio>Handler` (ej: `AuthHandler`).
- Constructores: `New<Dominio>Handler(...)`.
- Archivo: `<dominio>_handler.go` en snake_case.

## Ejemplo

```go
type AuthHandler struct {
    authService port.AuthServicePort
}

func NewAuthHandler(authService port.AuthServicePort) *AuthHandler { ... }
func (h *AuthHandler) Login(c *gin.Context) { ... }
```
