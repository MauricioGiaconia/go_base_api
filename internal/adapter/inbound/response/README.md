# Response

Contiene los **modelos de respuesta HTTP** y sus constructores.

## Responsabilidades

- Definir las estructuras de respuesta de la API (`SuccessResponse`, `SuccessListResponse`, `ErrorResponse`).
- Proveer la función `ToApi()` para construir respuestas estandarizadas.
- Proveer constructores de error (`NewError()`).

## Reglas

- Estos modelos son **artefactos de transporte**, NO entidades de dominio.
- Pueden tener JSON tags ya que representan el formato de comunicación HTTP.
- Nomenclatura: `<Tipo>Response` (ej: `SuccessResponse`, `ErrorResponse`).
- Función principal: `ToApi(code, data, isAList, count, limit, offset)`.
- Archivo: `api_response.go`.
