package error

import (
	"errors"
	"strings"
)

var (
	ErrInvalidCity      = errors.New("invalid city")
	ErrInvalidDirection = errors.New("invalid direction")
	ErrInvalidFileName  = errors.New("invalid filename")
)

func Wrap(err error, description string) error {
	return errors.New(strings.Join([]string{err.Error(), description}, " : "))
}
