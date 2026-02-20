# Application

Capa de **casos de uso** de la arquitectura hexagonal.

## Responsabilidades

- Orquestar la lógica de negocio utilizando los puertos definidos en `domain/port/`.
- Implementar los **input ports** (puertos de entrada).
- Consumir las dependencias externas únicamente a través de **output ports** (puertos de salida).

## Reglas

- Depende solo de `domain/` (entidades y puertos).
- No importar paquetes de infraestructura directamente (`gin`, `sql`, `jwt`, etc.).
- Las dependencias se inyectan a través de los constructores (`New...Service`).

## Subcarpetas

- `service/` — Implementaciones de los casos de uso.
