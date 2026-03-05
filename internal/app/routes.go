package app

import (
	"go-api-project/internal/features/users"
	userRepository "go-api-project/internal/features/users/repository"
	userService "go-api-project/internal/features/users/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a *App) registerRoutes() http.Handler {

	r := gin.Default()

	r.Use(cors.New(a.corsConfig))
	

	userRepo := userRepository.NewMongoUserRepository(a.db)

	userController := users.NewUserController(userService.NewUserService(userRepo))

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
