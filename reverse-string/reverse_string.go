// package reverse ...
package reverse

import (
	"unicode/utf8"
)

// String ...
func String(str string) string {
	if str == "" {
		return str
	}

	// convert string to a rune slice.
	var rs []rune
	for _, s := range str {
		r, _ := utf8.DecodeRuneInString(string(s))
		rs = append(rs, r)
	}

	max := len(rs) - 1
	var reversed []rune
	for i, _ := range rs {
		reversed = append(reversed, rs[max-i])
	}

	return string(reversed)
}
