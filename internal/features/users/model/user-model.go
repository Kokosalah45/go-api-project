package model

type User struct {
	ID          int     `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	IsAdmin     bool    `json:"is_admin"`
	Description *string `json:"description,omitempty"`
}
