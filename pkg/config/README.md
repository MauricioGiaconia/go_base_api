# Config

Utilidades de configuración **transversales** al sistema.

## Responsabilidades

- Cargar variables de entorno desde archivos `.env`.
- Proveer funciones de configuración utilizadas por múltiples capas (adapters, main).

## Reglas

- No contiene lógica de negocio.
- Es un paquete compartido (`pkg/`) porque es utilizado tanto por adapters inbound como outbound y por el punto de entrada.
- Nomenclatura: funciones utilitarias con nombres descriptivos (ej: `LoadEnv`).
- Archivo: `environment.go`.
