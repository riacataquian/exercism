// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Remove dashes.
	i := strings.Replace(s, "-", " ", -1)
	// Remove whitespaces.
	i = strings.Replace(i, " ", " ", -1)

	str := strings.Split(i, " ")

	var a string
	for _, s := range str {
		a += strings.ToUpper(s[:1])
	}

	return a
}
