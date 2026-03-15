package app

import (
	"go-api-project/bff-users/features/users/adapters/incoming/rest"
	"go-api-project/bff-users/features/users/adapters/outgoing/repository"
	"go-api-project/bff-users/features/users/service"
	"go-api-project/internal/logger"
	"go-api-project/internal/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a *App) registerRoutes() http.Handler {

	r := gin.New()

	r.Use(
		cors.New(a.corsConfig),
		gin.Recovery(),
		middleware.RequestIDMiddleware(),
		logger.LoggerMiddleware(a.logger),
	)

	userRepo := repository.NewMongoUserRepository(a.db)
	userService := service.NewUserService(userRepo)
	userController := rest.NewUserController(userService)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			userController.RegisterRoutes(v1)
		}
	}

	r.GET("/health", a.healthHandler)

	return r
}

func (a *App) healthHandler(c *gin.Context) {
	// Get logger from context and log health check
	log := logger.GetLoggerFromContext(c.Request.Context())
	log.Info("Health check requested", 
		logger.Str("request_id", middleware.GetRequestIDFromContext(c.Request.Context())))
	
	c.JSON(http.StatusOK, a.db.Health())
}
