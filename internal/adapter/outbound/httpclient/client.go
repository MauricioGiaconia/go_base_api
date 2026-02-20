package httpclient

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MauricioGiaconia/go_base_api/internal/adapter/inbound/response"
)

// ApiCall realiza una llamada HTTP externa con las opciones proporcionadas.
// Es un adaptador outbound que encapsula la comunicación con servicios externos.
func ApiCall(url string, opts ApiCallOptions) (interface{}, error) {

	// Método default en caso de que no se especifique
	if opts.Method == "" {
		opts.Method = "GET"
	}

	// Timeout por defecto en caso de que venga con un valor de 0
	if opts.Timeout == 0 {
		opts.Timeout = 5000
	}

	client := &http.Client{
		Timeout: time.Duration(opts.Timeout) * time.Millisecond,
	}

	req, err := http.NewRequest(opts.Method, url, bytes.NewBuffer(opts.Body))
	if err != nil {
		return response.ToApi(500, "Failed to create request", false, 0, 0, 0), err
	}

	req.Header.Set("access_token", "Bearer "+opts.Headers.AccessToken)
	if opts.Headers.ContentType != "" {
		req.Header.Set("Content-Type", opts.Headers.ContentType)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return response.ToApi(500, "Failed to execute request", false, 0, 0, 0), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.ToApi(500, "Failed to read response body", false, 0, 0, 0), err
	}

	if resp.StatusCode >= 400 {
		return response.ToApi(int64(resp.StatusCode), "API call failed", false, 0, 0, 0), errors.New(string(body))
	}

	return response.ToApi(200, string(body), false, 0, 0, 0), nil
}

// ApiCallWithFormat es una variante que permite especificar el formato de la respuesta.
func ApiCallWithFormat(url string, opts ApiCallOptions, isAList bool, count int64, limit int64, offset int64) (interface{}, error) {
	result, err := ApiCall(url, opts)
	if err != nil {
		return result, err
	}

	return response.ToApi(200, fmt.Sprintf("%v", result), isAList, count, limit, offset), nil
}
