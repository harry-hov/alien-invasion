package utils_test

import (
	"testing"

	"github.com/harry-hov/alien-invasion/utils"
	"github.com/harry-hov/alien-invasion/worldmap"
	"github.com/stretchr/testify/assert"
)

func TestPrettyJoinAliens(t *testing.T) {
	// Test aliens > 2
	aliens := []worldmap.Alien{"alien-0", "alien-1", "alien-2"}
	expected := "alien-0, alien-1, and alien-2"
	actual := utils.PrettyJoinAliens(aliens)
	assert.Equal(t, expected, actual)

	// Test aliens = 2
	aliens = []worldmap.Alien{"alien-0", "alien-1"}
	expected = "alien-0 and alien-1"
	actual = utils.PrettyJoinAliens(aliens)
	assert.Equal(t, expected, actual)
}
