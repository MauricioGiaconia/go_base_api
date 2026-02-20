# Service

Contiene las **implementaciones de los casos de uso** del sistema.

## Responsabilidades

- Implementar las interfaces definidas en `domain/port/` (input ports).
- Orquestar la lógica de negocio coordinando los output ports inyectados.

## Reglas

- Cada servicio recibe sus dependencias (output ports) por inyección en el constructor.
- Nomenclatura: `<Dominio>Service` (ej: `AuthService`).
- Constructores: `New<Dominio>Service(...)` (ej: `NewAuthService(...)`).
- Archivo: `<dominio>_service.go` en snake_case.

## Ejemplo

```go
type AuthService struct {
    userRepo     port.UserRepositoryPort
    tokenService port.TokenServicePort
}

func NewAuthService(userRepo port.UserRepositoryPort, tokenService port.TokenServicePort) *AuthService { ... }
```
