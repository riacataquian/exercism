// package luhn is a simple checksum formula used to validate a variety of identification numbers,
// such as credit card numbers and Canadian Social Insurance Numbers.
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid validates if input is a valid credit card number.
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

	// The first step of the Luhn algorithm is to double every second digit, starting from the right.
	in = reverse(in)

	var t int
	for i, v := range in {
		s := string(v)
		if s == "0" {
			continue
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			// Uhm, but dont' panic.
			panic(err)
		}

		// Check if we are to double the current iteration value.
		if i%2 == 0 {
			t += n
		} else {
			var num int
			// If doubling the number results in a number greater than 9 then subtract 9
			// from the product.
			num = n * 2
			if num > 9 {
				num -= 9
				t += num
			} else {
				t += num
			}
		}
	}

	// If the sum is evenly divisible by 10, then the number is valid.
	return t%10 == 0
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
