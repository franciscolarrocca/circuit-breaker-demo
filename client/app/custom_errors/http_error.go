package custom_errors

import "fmt"

var (
	Http4xxErrorCodes = map[int]struct{}{
		400: {},
		401: {},
		403: {},
		404: {},
		405: {},
	}
	Http5xxErrorCodes = map[int]struct{}{
		500: {},
		503: {},
	}
)

func Is5xxError(status int) bool {
	_, ok := Http5xxErrorCodes[status]
	return ok
}

func Is4xxError(status int) bool {
	_, ok := Http4xxErrorCodes[status]
	return ok
}

type HttpClientError struct {
	Status  int
	Message string
}

func (h *HttpClientError) Error() string {
	return fmt.Sprintf("an HTTP client error ocurred: '%d: %s'", h.Status, h.Message)
}

type HttpServerError struct {
	Status  int
	Message string
}

func (h *HttpServerError) Error() string {
	return fmt.Sprintf("an HTTP server error ocurred: '%d: %s'", h.Status, h.Message)
}
