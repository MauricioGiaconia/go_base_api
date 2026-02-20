# MySQL Adapter

Adaptador outbound que implementa los **output ports** de persistencia utilizando **MySQL**.

## Responsabilidades

- Implementar `UserRepositoryPort` → `UserRepository`.
- Implementar `ContractRepositoryPort` → `ContractRepository`.
- Proveer la conexión a la base de datos (`NewConnection()`).

## Reglas

- Cada repositorio implementa una interfaz definida en `domain/port/`.
- La conexión `*sql.DB` se inyecta en los constructores de los repositorios.
- Nomenclatura: `<Entidad>Repository` (ej: `UserRepository`).
- Constructores: `New<Entidad>Repository(db *sql.DB)`.
- Archivo: `<entidad>_repository.go` en snake_case.
- Conexión: `connection.go` con función `NewConnection()`.

## Nota de portabilidad

Si se necesita cambiar de MySQL a otro motor (PostgreSQL, MongoDB, etc.), basta con crear una nueva carpeta de adapter (ej: `outbound/postgres/`) que implemente los mismos ports.
