package users

import (
	"go-api-project/internal/features/users/dtos"
	"go-api-project/internal/features/users/service"
	"net/http"
	"strconv"

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
		users.POST("/", uc.CreateUser)
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var dto dtos.CreateUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userModel = dto.ToModel()

	user, err := uc.userService.CreateUser(c.Request.Context(), userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return	
	}

	c.JSON(http.StatusCreated, user)
}


func (uc *UserController) GetUserByID(c *gin.Context) {
	
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.userService.GetUserByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return	
	}

	c.JSON(http.StatusOK, user)

}