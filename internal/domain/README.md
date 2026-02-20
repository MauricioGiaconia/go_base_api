# Domain

Capa central de la arquitectura hexagonal. **No depende de ninguna otra capa**.

## Responsabilidades

- Definir las **entidades de negocio** (entities).
- Definir los **puertos** (interfaces) que describen las operaciones del sistema.

## Reglas

- No importar paquetes de infraestructura (`gin`, `sql`, `jwt`, etc.).
- No importar paquetes de los adapters (`inbound`, `outbound`).
- Solo Go puro y lógica de negocio.

## Subcarpetas

- `entity/` — Entidades de dominio.
- `port/` — Interfaces (puertos de entrada y salida).
