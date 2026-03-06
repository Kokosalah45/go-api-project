package main

import (
	"log"
	"net/http"

	"go-api-project/internal/app"
)

func main() {
	a := app.NewApp()

	done := make(chan bool, 1)

	go a.GracefulShutdown(done)

	if err := a.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}

	<-done
}
