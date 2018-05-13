// package bob ...
package bob

import (
	"fmt"
	"strings"
	"unicode"
)

// Hey ...
func Hey(remark string) string {
	yells := isYelling(remark)
	asks := isAsking(remark)

	if yells && asks {
		fmt.Printf("YELLING & ASKING %s\n", remark)
		return "Calm down, I know what I'm doing!"
	}

	if yells {
		fmt.Printf("YELLING %s\n", remark)
		return "Whoa, chill out!"
	}

	if asks {
		fmt.Printf("ASKING %s\n", remark)
		return "Sure."
	}

	fmt.Printf("MEH %s\n", remark)
	return "Whatever."
}

func isYelling(in string) bool {
	in = strimSpace(in)
	isYelling := true
	hasNumbers := false

	for _, l := range in {
		if unicode.IsDigit(l) {
			hasNumbers = true
		}

		if unicode.IsLower(l) && string(l) != "!" && string(l) != "?" {
			isYelling = false
		}
	}

	if hasNumbers {
		isYelling = strings.Contains(in, "!")
	}

	return isYelling
}

func isAsking(in string) bool {
	return strings.HasSuffix(in, "?")
}

func strimSpace(in string) string {
	return strings.Replace(in, " ", "", -1)
}
