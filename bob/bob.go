// package bob ...
package bob

import (
	"strings"
	"unicode"
)

// Hey ...
func Hey(remark string) string {
	in := stripWhitespaces(remark)
	yells := isYelling(in)
	asks := isAsking(in)

	if yells && asks {
		return "Calm down, I know what I'm doing!"
	}

	if yells {
		return "Whoa, chill out!"
	}

	if asks {
		return "Sure."
	}

	if in == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func stripWhitespaces(remark string) string {
	str := strings.Replace(remark, " ", "", -1)
	str = strings.Trim(str, " ")
	str = strings.Trim(str, "\t")
	str = strings.Trim(str, "\r")
	str = strings.Trim(str, "\n")

	return str
}

func isYelling(in string) (isYelling bool) {
	hasLowerCases := false

	if strings.HasSuffix(in, ".") {
		isYelling = false
		return
	}

	chars := strings.Replace(in, "!", "", -1)
	chars = strings.Replace(chars, ".", "", -1)
	for _, r := range chars {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				isYelling = true
			} else {
				hasLowerCases = true
			}
		}
	}

	if hasLowerCases {
		isYelling = false
		return
	}

	return
}

func isAsking(in string) bool {
	return strings.HasSuffix(in, "?")
}
