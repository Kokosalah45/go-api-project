package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-api-project/internal/config"
	"go-api-project/internal/database"
	"go-api-project/internal/logger"
	"go-api-project/internal/database/mongodb"

	"github.com/gin-contrib/cors"
)

type App struct {
	port       int
	db         database.DBService
	server     *http.Server
	corsConfig cors.Config
	logger     logger.Logger
}

func NewApp() *App {
	
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)                           
	}

	// Initialize logger with configuration
	baseLogger := logger.NewLoggerFromConfig(logger.Config{
		Level:  cfg.Log.Level,
		Format: cfg.Log.Format,
	})
	baseLogger.Info("Application starting up", 
		logger.Str("app_env", cfg.App.AppEnv),
		logger.Int("port", cfg.App.Port))

	db, err := mongodb.New(&mongodb.MongoDBConf{
		Host:    cfg.DB.Host,
		Port:    cfg.DB.Port,
		AppUser: cfg.DB.AppUser,
		AppPass: cfg.DB.AppPass,
		DBName:  cfg.DB.DBName,
	})

	if err != nil {
		baseLogger.Fatal("Failed to connect to MongoDB", logger.Err(err))
	}
	
	baseLogger.Info("Database connection established")

	newApp := &App{
		port: cfg.App.Port,
		db:   db,
		logger: baseLogger,
		corsConfig: cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
		},
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.App.Port),
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

	a.logger.Info("Shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.Shutdown(ctx); err != nil {
		a.logger.Error("Server forced to shutdown", logger.Err(err))
	}

	a.logger.Info("Server exiting")
	done <- true
}

func (a *App) Start() error {
	a.logger.Info("Server starting", logger.Int("port", a.port))
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	a.logger.Info("Shutting down server...")
	if err := a.db.Close(); err != nil {
		a.logger.Error("Error closing DB", logger.Err(err))
	}
	return a.server.Shutdown(ctx)
}
