// package bob handles and processes responses by Bob.
package bob

import (
	"strings"
	"unicode"
)

// DefaultResponse is Bob's default response if remark is not asking, yelling or is silent.
const DefaultResponse = "Whatever."

// Hey returns response relative to Bob's remark.
func Hey(remark string) string {
	in := stripWhitespaces(remark)
	if strings.HasSuffix(in, ".") {
		return DefaultResponse
	}

	yells := isYelling(in)
	asks := isAsking(in)
	silence := in == ""

	if yells && asks {
		return "Calm down, I know what I'm doing!"
	}

	if yells {
		return "Whoa, chill out!"
	}

	if asks {
		return "Sure."
	}

	if silence {
		return "Fine. Be that way!"
	}

	return DefaultResponse
}

func stripWhitespaces(remark string) string {
	return strings.TrimFunc(remark, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}

// isYelling loops through the input string then scans for lowercase letters,
// if found, considered `in` as not yelling.
// func isYelling(in string) (isYelling bool) {
func isYelling(in string) (isYelling bool) {
	for _, r := range in {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				isYelling = true
			} else {
				isYelling = false
				break
			}
		}
	}

	return
}

// isAsking considers inputs with "?" as suffix to be asking.
func isAsking(in string) bool {
	return strings.HasSuffix(in, "?")
}
