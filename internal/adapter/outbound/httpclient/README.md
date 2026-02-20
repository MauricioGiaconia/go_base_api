# HTTP Client Adapter

Adaptador outbound para realizar **llamadas HTTP a APIs externas**.

## Responsabilidades

- Proveer una función `ApiCall()` para realizar peticiones HTTP con configuración flexible.
- Definir los tipos de opciones y headers para las llamadas (`ApiCallOptions`, `Headers`).

## Reglas

- Las respuestas utilizan los modelos de `adapter/inbound/response/` para mantener consistencia.
- Nomenclatura de tipos: `ApiCallOptions`, `Headers`.
- Función principal: `ApiCall(url, opts)`.
- Archivos: `types.go` (modelos), `client.go` (lógica de llamadas).
