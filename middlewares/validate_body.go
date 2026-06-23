package middlewares

import (
	"Authingo/utils"
	"context"
	"net/http"
)

// custom type for context key, taaki collision na ho doosre packages ke keys se
type contextKey string

const PayloadContextKey contextKey = "validatedPayload"



func ValidateBody[T any](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload T

		if err := utils.ReadJson(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "invalid request body", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(), PayloadContextKey, payload)
		next(w, r.WithContext(ctx))
	}
}