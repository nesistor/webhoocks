# Wybierz obraz bazowy z oficjalnej wersji Go
FROM golang:1.21-alpine as builder

# Ustaw katalog roboczy w kontenerze
WORKDIR /app

# Skopiuj pliki go.mod i go.sum, aby zainstalować zależności
COPY go.mod go.sum ./
RUN go mod tidy

# Skopiuj całą resztę aplikacji do kontenera
COPY . .

# Zbuduj aplikację w trybie produkcyjnym
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Użyj minimalnego obrazu do uruchomienia aplikacji
FROM alpine:latest

# Instaluj całą niezbędną konfigurację w Alpine (np. certyfikaty SSL, aby wspierać HTTPS)
RUN apk --no-cache add ca-certificates

# Skopiuj skompilowaną aplikację z etapu build
COPY --from=builder /app/webhoockapp /webhoockapp

# Expose port aplikacji
EXPOSE 8080

# Komenda uruchamiająca aplikację
CMD ["/webhoockapp"]
