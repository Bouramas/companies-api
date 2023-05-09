package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var (
	ContextTokenKey = contextKey("token")
)

type contextKey string

func AuthHandler() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				writeJSONUnauthorizedError(w, "Unauthorized access - missing Authorization header")
				return
			}

			schemeToken := strings.Split(authHeader, " ")
			if len(schemeToken) != 2 {
				writeJSONUnauthorizedError(w, "Invalid authorization header.")
				return
			}

			scheme := strings.ToLower(schemeToken[0])
			if scheme != "bearer" {
				writeJSONUnauthorizedError(w, "Invalid authorization header scheme.")
				return
			}

			jwtToken := schemeToken[1]
			if jwtToken == "" {
				writeJSONUnauthorizedError(w, "No authorization token provided.")
				return
			}

			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte("my_secret_key"), nil
			})

			if !token.Valid {
				writeJSONUnauthorizedError(w, "Unauthorized access - Invalid Token")
				return
			}

			ctx := context.WithValue(r.Context(), ContextTokenKey, token)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func writeJSONUnauthorizedError(w http.ResponseWriter, msg string) {
	errObject := map[string]interface{}{"error": true, "code": http.StatusUnauthorized, "message": msg}
	res, _ := json.Marshal(errObject)
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write(res)
}
