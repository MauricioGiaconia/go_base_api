package httpclient

// ApiCallOptions define las opciones de configuración para realizar llamadas HTTP externas.
type ApiCallOptions struct {
	Headers Headers
	Method  string
	Timeout int64
	Body    []byte
}

// Headers define los headers HTTP que se enviarán en las llamadas externas.
type Headers struct {
	AccessToken string
	ContentType string
}
