// package hamming measures the minimum number of point mutations
// that could have occurred on the evolutionary path between two strands.
package hamming

import (
	"errors"
)

// Distance calcualtes difference between a and b stands.
func Distance(a, b string) (int, error) {
	// Prevent unmatched strand lengths
	if len(a) != len(b) {
		return -1, errors.New("Mismatched lengths")
	}

	// If strands matched, then return 0 diff.
	if a == b {
		return 0, nil
	}

	// Compare per letter, accumulate diff if unmatched string value.
	var diff int
	for i, l := range a {
		if string(l) != string(b[i]) {
			diff++
		}
	}
	return diff, nil
}
