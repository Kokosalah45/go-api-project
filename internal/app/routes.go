package app

import (
	"go-api-project/internal/features/users/adapters/incoming/rest"
	"go-api-project/internal/features/users/adapters/outgoing/repository"
	"go-api-project/internal/features/users/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a *App) registerRoutes() http.Handler {

	r := gin.New()

	r.Use(cors.New(a.corsConfig))
	r.Use(gin.Recovery())

	userRepo := repository.NewMongoUserRepository(a.db)
	userController := rest.NewUserController(service.NewUserService(userRepo))

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
	c.JSON(http.StatusOK, a.db.Health())
}
