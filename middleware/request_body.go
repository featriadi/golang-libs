package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	responsebuilder "github.com/featriadi/golang-libs/response_builder"
	"github.com/featriadi/golang-libs/validator"
)

func RequestBodyValidator[T any]() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var requestBody T

			if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
				errMsg := err.Error()
				responsebuilder.NewResponseBuilder[string](&w).
					Status(http.StatusBadRequest).
					Message("Invalid request body").
					Errors(&errMsg).
					Build()
				return
			}

			if err := validator.Validator(requestBody); err != nil {
				responsebuilder.NewResponseBuilder[map[string]string](&w).
					Status(http.StatusBadRequest).
					Message("Validation error").
					Errors(&err).
					Build()
				return
			}

			ctx := context.WithValue(r.Context(), requestBodyKey, requestBody)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
