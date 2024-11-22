package main

import (
	"fmt"
	"net/http"
	"strings"
)

var globalToken = ""

// GenerateTokenMiddleware generuje token globalny i wypisuje go w konsoli
func GenerateTokenMiddleware() {
	// Przykładowy statyczny token - możesz zmienić na bardziej dynamiczny mechanizm
	globalToken = "super_secure_global_token_12345"
	fmt.Printf("Generated Token: %s\n", globalToken)
}

// ValidateTokenMiddleware sprawdza poprawność tokena w nagłówku Authorization
func ValidateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pobierz nagłówek Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Sprawdź, czy nagłówek ma format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized: invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Sprawdź, czy token jest zgodny z globalnym tokenem
		token := parts[1]
		if token != globalToken {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// Przekaż żądanie do następnego handlera
		next.ServeHTTP(w, r)
	})
}
