package rest

import (
	"go-api-project/bff-users/features/common"
	"go-api-project/bff-users/features/users/adapters/incoming/rest/dtos"
	"go-api-project/bff-users/features/users/service"
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

func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", uc.CreateUser)
		users.GET("/:id", uc.GetUserByID)
		users.PUT("/:id", uc.UpdateUser)
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var dto dtos.CreateUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.NewBadRequestError()})
		return
	}

	var userModel = dto.ToModel()

	user, err := uc.userService.CreateUser(c.Request.Context(), userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUserByID(c *gin.Context) {

	idParam := c.Param("id")

	user, err := uc.userService.GetUserByID(c.Request.Context(), idParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	userResponse := dtos.NewDetailedUserResponse(user)

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserController) UpdateUser(c *gin.Context) {

	idParam := c.Param("id")

	var dto dtos.UpdateUserRequest

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.NewBadRequestError()})
		return
	}

	updatedUser, err := uc.userService.UpdateUser(c.Request.Context(), idParam, &dto)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	userResponse := dtos.NewDetailedUserResponse(updatedUser)

	c.JSON(http.StatusOK, userResponse)
}
