package domain

type User struct {
	ID          string
	Username    string
	Email       string
	Description *string
	Age         *int
}

func NewUser(username, email string, description *string, age *int) *User {
	return &User{
		Username:    username,
		Email:       email,
		Description: description,
		Age:         age,
	}
}
