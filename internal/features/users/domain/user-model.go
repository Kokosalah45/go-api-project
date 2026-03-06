package domain

type User struct {
	ID          string
	Username    string
	Email       string
	Description *string
	Age         *int
}

func (u *User) AddDescription(desc *string) error {

	if u.Description != nil {
		return NewUserHasAlreadyDescriptionError()
	}

	u.Description = desc

	return nil
}
