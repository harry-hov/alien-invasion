package error_test

import (
	"errors"
	"testing"

	e "github.com/harry-hov/alien-invasion/error"
	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	actual := e.Wrap(e.ErrInvalidCity, "invalid city")
	expected := errors.New("invalid city : invalid city")
	assert.Equal(t, expected, actual)
}
