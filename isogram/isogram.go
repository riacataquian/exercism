// Package isogram ...
package isogram

import (
	"strings"
	"unicode"
)

var charset map[string]bool

// IsIsogram ...
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
