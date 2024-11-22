package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"log"
	"net/http"
	"strings"
)

// StaticToken przechowuje statyczny token
var StaticToken string

// GenerateStaticToken generuje jeden losowy token statyczny
func GenerateStaticToken() string {
	// Generate 16 random bytes
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("Error generating token: %v", err)
	}

	// Encode the random bytes to a base32 string
	token := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// Opcjonalne: wyświetl hash tokena
	hash := sha256.Sum256([]byte(token))
	log.Printf("Generated Token: %s", token)
	log.Printf("Token Hash: %x", hash)

	return token
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
		if token != StaticToken {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// Przekaż żądanie do następnego handlera
		next.ServeHTTP(w, r)
	})
}
	