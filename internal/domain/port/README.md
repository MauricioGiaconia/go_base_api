# Port

Contiene las **interfaces** (puertos) que definen los contratos entre capas.

## Tipos de puertos

### Input Ports (puertos de entrada)

Interfaces que los **adaptadores inbound** (handlers HTTP) invocan para ejecutar la lógica de negocio.

- Nomenclatura: `<Dominio>ServicePort` (ej: `AuthServicePort`).
- Implementados por los services en `application/service/`.

### Output Ports (puertos de salida)

Interfaces que la **capa de aplicación** utiliza para interactuar con el mundo externo (BD, APIs, etc.)

- Nomenclatura: `<Entidad>RepositoryPort` (ej: `UserRepositoryPort`, `ContractRepositoryPort`).
- Nomenclatura servicios: `<Concepto>ServicePort` (ej: `TokenServicePort`).
- Implementados por los adapters en `adapter/outbound/`.

## Reglas

- Solo interfaces, nunca structs concretas.
- No importar paquetes de infraestructura.
- Solo pueden importar del paquete `entity`.
