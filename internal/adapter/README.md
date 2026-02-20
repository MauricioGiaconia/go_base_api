# Adapter

Capa de **adaptadores** de la arquitectura hexagonal. Conecta el mundo externo con el dominio.

## Subcarpetas

### `inbound/` (adaptadores de entrada / driving adapters)

Reciben peticiones del exterior y las traducen a llamadas a los input ports.

- `handler/` — Controladores HTTP (Gin handlers).
- `middleware/` — Middlewares HTTP (autenticación JWT, validación de contratos).
- `router/` — Configuración de rutas de la API.
- `response/` — Modelos de respuesta HTTP y constructores de respuestas.

### `outbound/` (adaptadores de salida / driven adapters)

Implementan los output ports para interactuar con servicios externos.

- `mysql/` — Repositorios MySQL (implementan `UserRepositoryPort`, `ContractRepositoryPort`).
- `httpclient/` — Cliente HTTP para llamadas a APIs externas.

## Reglas

- Los adapters inbound dependen de los **input ports** (`domain/port/`).
- Los adapters outbound implementan los **output ports** (`domain/port/`).
- Pueden importar librerías de infraestructura (`gin`, `sql`, `jwt`, etc.).
