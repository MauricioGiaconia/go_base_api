# Entity

Contiene las **entidades de dominio** del sistema.

## Responsabilidades

- Representar los conceptos de negocio con sus atributos.
- Pueden contener lógica de validación propia de la entidad.

## Reglas

- No deben tener dependencias de infraestructura.
- No deben contener JSON tags de respuesta HTTP (esos van en `adapter/inbound/response/`).
- Nomenclatura: nombre del concepto en singular y PascalCase (ej: `User`, `Contract`).

## Ejemplos

- `user.go` → Entidad `User`
- `contract.go` → Entidad `Contract`
