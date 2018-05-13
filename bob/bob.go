// package bob ...
package bob

import (
	"strings"
	"unicode"
)

// Hey ...
func Hey(remark string) string {
	if isYelling(remark) {
		return "Whoa, chill out!"
	}

	if isAsking(remark) {
		return "Sure."
	}

	return "Whatever."
}

func isYelling(in string) bool {
	in = strimSpace(in)
	isYelling := true

	for _, l := range in {
		if !unicode.IsUpper(l) && string(l) != "!" {
			isYelling = false
		}
	}

	return isYelling
}

func isAsking(in string) bool {
	in = strimSpace(in)
	isAsking := true

	for _, l := range in {
		if !unicode.IsUpper(l) && string(l) != "?" {
			isAsking = false
		}
	}

	return isAsking
}

func strimSpace(in string) string {
	return strings.Replace(in, " ", "", -1)
}
