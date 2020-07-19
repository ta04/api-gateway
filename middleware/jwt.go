package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ta04/api-gateway/internal/config"
	authModel "github.com/ta04/auth-service/model"
)

// Error is an error message
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Response is a response message
type Response struct {
	Error *Error `json:"error,omitempty"`
}

// JWTMiddleware is the middleware for JWT based auth
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unauthorizedRes := &Response{
			Error: &Error{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized Access. Malformed Token",
			},
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			marshaledRes, err := json.Marshal(unauthorizedRes)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(marshaledRes)
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.ParseWithClaims(jwtToken, &authModel.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.SecretKey()), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			claims, ok := token.Claims.(*authModel.Claims)
			if ok && token.Valid {
				ctx := context.WithValue(r.Context(), "claims", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				marshaledRes, err := json.Marshal(unauthorizedRes)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusUnauthorized)
				w.Write(marshaledRes)
			}
		}
	})
}
