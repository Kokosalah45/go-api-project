package dtos

type UpdateUserRequest struct {
	Username    *string `json:"username" binding:"required"`
	Email       *string `json:"email" binding:"required,email"`
	Description *string `json:"description" binding:"omitempty,min=25,max=500"`
	Age         *int    `json:"age" binding:"omitempty,gte=25,lte=55"`
}
