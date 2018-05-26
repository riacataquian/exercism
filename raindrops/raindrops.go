// Package raindrops converts a number to a string,
// the contents of which depend on the number's factors.
// http://exercism.io/exercises/go/raindrops/readme
package raindrops

import (
	"strconv"
)

// kv maps factor to its speak equivalent.
type kv struct {
	factor int
	speak  string
}

var lookup = []kv{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert calculates for the number's factor and
// returns the appropriate raindrop speak.
func Convert(n int) string {
	var speak string
	for _, l := range lookup {
		if n == l.factor {
			speak += l.speak
			break
		}

		if n%l.factor == 0 {
			speak += l.speak
		}
	}

	if speak != "" {
		return speak
	}

	// If no factors were found, return the the int as string.
	return strconv.Itoa(n)
}
