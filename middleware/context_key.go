package middleware

import "net/http"

type contextKey string

const requestBodyKey contextKey = "requestBody"

func GetRequestBody[T any](r *http.Request) *T {
	data, ok := r.Context().Value(requestBodyKey).(T)
	if !ok {
		return nil
	}

	return &data
}
