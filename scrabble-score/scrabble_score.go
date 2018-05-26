// Package scrabble given a word, compute the scrabble score for that word.
package scrabble

import (
	"strings"
)

// Score performs the calculation of the score per character of a given word.
func Score(word string) int {
	ref := buildLookup()
	str := strings.ToUpper(word)

	// nit: optimization
	if len(str) == 1 {
		return ref[str]
	}

	var score int
	for _, r := range str {
		// r is in rune type, cast to string.
		score += ref[string(r)]
	}

	return score
}

// buildLookup prepares and build the lookup table of alphabets and their scrabled score equivalents.
func buildLookup() map[string]int {
	var lookup = make(map[string]int)
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
	return lookup
}
