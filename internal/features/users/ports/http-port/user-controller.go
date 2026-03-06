package httpport

import (
	"fmt"
	"go-api-project/internal/features/common"
	"go-api-project/internal/features/users/ports/http-port/dtos"
	"go-api-project/internal/features/users/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.POST("", uc.CreateUser)
		users.GET("/:id", uc.GetUserByID)
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var dto dtos.CreateUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.NewBadRequestError()})
		return
	}

	var userModel = dto.ToModel()

	fmt.Println(userModel.Age,dto.Age)

	user, err := uc.userService.CreateUser(c.Request.Context(), userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUserByID(c *gin.Context) {

	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.userService.GetUserByID(c.Request.Context(), idParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	userResponse := dtos.NewDetailedUserResponse(user)

	c.JSON(http.StatusOK, userResponse)
}
