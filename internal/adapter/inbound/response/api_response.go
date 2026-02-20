package response

import "fmt"

// SuccessResponse representa una respuesta exitosa de la API para un único recurso.
type SuccessResponse struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
}

// SuccessListResponse representa una respuesta exitosa de la API para un listado con paginación.
type SuccessListResponse struct {
	Code     int64       `json:"code"`
	Data     interface{} `json:"data"`
	Count    int64       `json:"count"`
	Limit    int64       `json:"limit"`
	Offset   int64       `json:"offset"`
	Next     string      `json:"next,omitempty"`
	Previous string      `json:"previous,omitempty"`
}

// ErrorResponse representa una respuesta de error de la API.
type ErrorResponse struct {
	Code  int64  `json:"code"`
	Error string `json:"error"`
}

// ToApi construye la respuesta adecuada según el código de estado y tipo de respuesta.
// Para una respuesta de un único dato exitoso, bastará con enviar "code" y "data".
// Para un listado con paginado, se deben enviar todos los argumentos.
func ToApi(code int64, data interface{}, isAList bool, count int64, limit int64, offset int64) interface{} {
	if code >= 400 {
		return buildErrorResponse(code, data)
	}

	if isAList {
		return buildSuccessListResponse(code, data, count, limit, offset)
	}

	return SuccessResponse{
		Code: code,
		Data: data,
	}
}

// NewError crea una respuesta de error con código y mensaje personalizados.
func NewError(code int64, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:  code,
		Error: message,
	}
}

func buildErrorResponse(code int64, data interface{}) ErrorResponse {
	if str, ok := data.(string); ok {
		return ErrorResponse{
			Code:  code,
			Error: str,
		}
	}
	return ErrorResponse{
		Code:  code,
		Error: "Sorry, we had a trouble processing the information",
	}
}

func buildSuccessListResponse(code int64, data interface{}, count int64, limit int64, offset int64) interface{} {
	return SuccessListResponse{
		Code:     code,
		Data:     data,
		Count:    count,
		Limit:    limit,
		Offset:   offset,
		Next:     calculateNextURL(limit, offset, count),
		Previous: calculatePreviousURL(limit, offset),
	}
}

func calculateNextURL(limit, offset, count int64) string {
	if offset+limit >= count {
		return ""
	}
	return fmt.Sprintf("?limit=%d&offset=%d", limit, offset+limit)
}

func calculatePreviousURL(limit, offset int64) string {
	if offset-limit < 0 {
		return ""
	}
	return fmt.Sprintf("?limit=%d&offset=%d", limit, offset-limit)
}
