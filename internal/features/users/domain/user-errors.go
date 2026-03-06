package domain

type UserHasAlreadyDiscriptionError struct {
	Code    string
	Message string
}

const (
	UserAlreadyHasDescriptionErrorCode    = "USER_ALREADY_HAS_DESCRIPTION"
	UserAlreadyHasDescriptionErrorMessage = "user already has a description"
)

func NewUserHasAlreadyDescriptionError() *UserHasAlreadyDiscriptionError {
	return &UserHasAlreadyDiscriptionError{
		Code:    UserAlreadyHasDescriptionErrorCode,
		Message: UserAlreadyHasDescriptionErrorMessage,
	}
}

func (e *UserHasAlreadyDiscriptionError) Error() string {
	return e.Message
}
