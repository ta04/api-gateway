package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ta04/api-gateway/config"
)

// Error is an error message
type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Response is a response message
type Response struct {
	Error *Error `json:"error,omitempty"`
}

// JWTMiddleware is the middleware for JWT based auth
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			res := &Response{
				Error: &Error{
					Code:    strconv.Itoa(http.StatusUnauthorized),
					Message: "Unauthorized Access. Malformed Token",
				},
			}

			marshaledRes, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(marshaledRes)
		} else {
			jwtToken := authHeader[1]
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.SecretKey()), nil
			})

			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				ctx := context.WithValue(r.Context(), "claims", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				res := &Response{
					Error: &Error{
						Code:    strconv.Itoa(http.StatusUnauthorized),
						Message: "Unauthorized Access",
					},
				}

				marshaledRes, err := json.Marshal(res)
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
