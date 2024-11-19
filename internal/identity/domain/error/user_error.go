package identity

import "errors"

var (
	ErrUserAlreadyExists = errors.New("error_user_already_exists")
)

func NewErrUserAlreadyExists() error {
	return ErrUserAlreadyExists
}
