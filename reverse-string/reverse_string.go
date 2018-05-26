// package reverse a given string.
package reverse

// String converts str to rune slices for later reversing.
func String(str string) string {
	if str == "" {
		return str
	}

	// convert string to a rune slice.
	rs := []rune(str)
	max := len(rs) - 1

	var reversed []rune
	for i, _ := range rs {
		reversed = append(reversed, rs[max-i])
	}

	return string(reversed)
}
