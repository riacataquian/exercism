// package luhn ...
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid ...
func Valid(in string) bool {
	// Strings of length 1 or less are not valid.
	if len(in) <= 1 {
		return false
	}

	isValid := true

	// Modify the input with the following criteria:
	// Spaces are allowed in the input, but they should be stripped before checking.
	in = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		if !unicode.IsDigit(r) {
			isValid = false
			return -1
		}

		return r
	}, in)

	// All other non-digit characters are disallowed, makes the input invalid.
	if !isValid {
		return isValid
	}

	// Zeroes are invalid.
	if in == "0" {
		return false
	}

	var t int
	for i, v := range in {
		s := string(v)
		n, err := strconv.Atoi(s)
		if err != nil {
			// Uhm, but dont' panic.
			panic(err)
		}

		// Check if we are to double the current iteration value.
		if i%2 == 0 {
			t += n
		} else {
			// If doubling the number results in a number greater than 9 then subtract 9
			// from the product.
			var num int
			num = n * 2
			if num > 9 {
				num -= 9
			}

			t += num
		}

	}

	return t%10 == 0
}
