// package twofer or 2-fer is short for two for one.
//
// See example_two_fer_test.go for example usage.
package twofer

// ShareWith returns One for X, one for me, where X is a name or "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
