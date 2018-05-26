// Package isogram encapsulate functions and helpers for dealing with isograms.
package isogram

import (
	"strings"
	"unicode"
)

// charset holds the existing characters.
var charset map[string]bool

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	if word == "" {
		return true
	}

	isIsogram := true
	charset = make(map[string]bool)
	for _, r := range word {
		// ignore non-letters.
		if !unicode.IsLetter(r) {
			continue
		}

		s := strings.ToUpper(string(r)) // don't mind the case!

		if _, ok := charset[s]; ok {
			isIsogram = false
			break
		}
		charset[s] = true
	}

	return isIsogram
}
