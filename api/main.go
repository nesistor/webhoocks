package main

import (
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct{}

func main() {

	StaticToken = GenerateStaticToken()
	log.Printf("Static Token: %s", StaticToken)

	app := Config{}
	log.Printf("Starting server on port %s\n", webPort)

	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
