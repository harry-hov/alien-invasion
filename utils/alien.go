package utils

import (
	"fmt"

	"github.com/harry-hov/alien-invasion/worldmap"
)

// PrettyJoinAliens join the list of alien in human
// readable format.
//
// e.g:
//
// Aliens["alien-1", "alien-2", "alien-3"] => "alien-1, alien-2 and alien-3"
//
// Aliens["alien-1", "alien-2"] => "alien-1 and alien-2"
func PrettyJoinAliens(aliens []worldmap.Alien) (out string) {
	n := len(aliens)
	switch n {
	case 2:
		out += fmt.Sprintf("%v ", aliens[0])
	default:
		for i := 0; i < n-1; i++ {
			out += fmt.Sprintf("%v, ", aliens[i])
		}
	}

	out += fmt.Sprintf("and %v", aliens[n-1])
	return
}
