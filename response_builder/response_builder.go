package responsebuilder

import (
	"encoding/json"
	"net/http"
)

type ResponseBuilder[T any] struct {
	w        *http.ResponseWriter
	response *ResponseJson[T]
}

func NewResponseBuilder[T any](w *http.ResponseWriter) *ResponseBuilder[T] {
	return &ResponseBuilder[T]{
		w:        w,
		response: NewResponseJson[T](),
	}
}

func (rb *ResponseBuilder[T]) Status(status int) *ResponseBuilder[T] {
	rb.response.SetStatus(status)
	return rb
}

func (rb *ResponseBuilder[T]) Message(message string) *ResponseBuilder[T] {
	rb.response.SetMessage(message)
	return rb
}

func (rb *ResponseBuilder[T]) Data(data *T) *ResponseBuilder[T] {
	rb.response.SetData(data)
	return rb
}

func (rb *ResponseBuilder[T]) Pagination(totalRows int, page int, size int) *ResponseBuilder[T] {
	rb.response.SetPagination(totalRows, page, size)
	return rb
}

func (rb *ResponseBuilder[T]) Errors(errors *T) *ResponseBuilder[T] {
	rb.response.SetErrors(errors)
	return rb
}

func (rb *ResponseBuilder[T]) Build() {
	(*rb.w).Header().Set("Content-Type", "application/json")
	(*rb.w).WriteHeader(rb.response.Status)

	encoder := json.NewEncoder(*rb.w)
	encoder.SetIndent("", "  ")
	encoder.Encode(rb.response)
}
