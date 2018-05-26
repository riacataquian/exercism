// Package scrabble ...
package scrabble

import (
	"strings"
	// "sync"
)

// Score ...
// TODO optimize for 1 chars
func Score(str string) int {
	ref := buildLookup()

	var score int
	for _, r := range str {
		// r is in rune type, cast to string.
		char := strings.ToUpper(string(r))
		score += ref[char]
	}

	return score
}

func buildLookup() map[string]int {
	var lookup = make(map[string]int)
	// var mutex = &sync.Mutex{}

	// mutex.Lock()
	lookup["A"] = 1
	lookup["B"] = 3
	lookup["C"] = 3
	lookup["D"] = 2
	lookup["E"] = 1
	lookup["F"] = 4
	lookup["G"] = 2
	lookup["H"] = 4
	lookup["I"] = 1
	lookup["J"] = 8
	lookup["K"] = 5
	lookup["L"] = 1
	lookup["M"] = 3
	lookup["N"] = 1
	lookup["O"] = 1
	lookup["P"] = 3
	lookup["Q"] = 10
	lookup["R"] = 1
	lookup["S"] = 1
	lookup["T"] = 1
	lookup["U"] = 1
	lookup["V"] = 4
	lookup["W"] = 4
	lookup["X"] = 8
	lookup["Y"] = 4
	lookup["Z"] = 10
	// mutex.Unlock()

	return lookup
}
