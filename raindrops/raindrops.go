// Package raindrops converst a number to a string,
// the contents of which depend on the number's factors.
// http://exercism.io/exercises/go/raindrops/readme
package raindrops

import (
	"strconv"
)

// rd describes a raindrop language.
var rd map[int]string

// Convert calculates for the number's factor and
// returns the appropriate raindrop speak.
func Convert(n int) string {
	rd = make(map[int]string)
	rd[3] = "Pling"
	rd[5] = "Plang"
	rd[7] = "Plong"

	var speak string
	for k, v := range rd {
		if n == k {
			speak += v
			break
		}

		if n%k == 0 {
			speak += v
		}
	}

	if speak != "" {
		return speak
	}

	// If no factors were found, return the the int as string.
	return strconv.Itoa(n)
}
