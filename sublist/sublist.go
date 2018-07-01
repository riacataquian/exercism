// Package sublist determines if a list is a sublist of another list.
package sublist

// Relation describes a relation between two lists.
type Relation string

const (
	// Unequal is when two lists are not equal.
	Unequal Relation = "unequal"

	// Equal is when two lists are equal.
	Equal Relation = "equal"

	// Sub is when a list is a sublist to another list.
	Sub Relation = "sublist"

	// Sup is when a list is a superlist to another list.
	Sup Relation = "superlist"
)

// Sublist returns the relation between two lists.
func Sublist(a, b []int) Relation {
	diff := len(a) - len(b)

	if diff == 0 && equal(a, b) {
		return Equal
	}

	if diff > 0 && sub(b, a) {
		return Sup
	}

	if diff < 0 && sub(a, b) {
		return Sub
	}

	return Unequal
}

// sub returns true if a is a sublist of b.
func sub(a, b []int) bool {
	// An empty list is always a sublist of another list.
	if len(a) == 0 {
		return true
	}

	var match []int
	isSub := false
loop:
	for i, m := range b {
		for _, n := range a {
			if m == n {
				max := i + len(a)

				// Traverse superlist from the current matching element's index
				// up to the sublist length.
				if max > len(b) {
					match = b[i:]
				} else {
					match = b[i:max]
				}

				isSub = equal(match, a)
			}
		}
		if isSub {
			break loop
		}
	}

	return isSub
}

// equal returns true if two lists are equal with the same order.
func equal(a, b []int) bool {
	for i, j := range a {
		if b[i] != j {
			return false
		}
	}
	return true
}
