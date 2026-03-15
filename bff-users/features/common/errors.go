package common

import "errors"

func NewBadRequestError() error {
	return errors.New("bad request")
}
