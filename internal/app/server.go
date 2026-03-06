package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"go-api-project/internal/database"
	mongodb "go-api-project/internal/database/mongodb"

	"github.com/gin-contrib/cors"
)

type App struct {
	port       int
	db         database.DBService
	server     *http.Server
	corsConfig cors.Config
}

func NewApp() *App {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	newApp := &App{
		port: port,
		db:   mongodb.New(),
		corsConfig: cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
		},
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}

	newApp.server.Handler = newApp.registerRoutes()

	return newApp
}

func (a *App) GracefulShutdown(done chan<- bool) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	fmt.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown with error: %v", err)
	}

	fmt.Println("Server exiting")
	done <- true
}

func (a *App) Start() error {
	fmt.Printf("Server is running on port %d\n", a.port)
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down server...")
	if err := a.db.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}
	return a.server.Shutdown(ctx)
}
