// Package sublist ...
package sublist

import (
// "log"
)

// Relation ...
type Relation string

const (
	Unequal Relation = "unequal"
	Equal   Relation = "equal"
	Sub     Relation = "sublist"
	Sup     Relation = "superlist"
	Wat     Relation = "wat"
)

func Sublist(a, b []int) Relation {
	if len(a) == 0 && len(b) == 0 {
		return Equal
	}

	if sub(a, b) {
		return Sub
	}

	if sup(b, a) {
		return Sup
	}

	if equal(a, b) {
		return Equal
	}

	return Unequal
}

func equal(a, b []int) bool {
	var matching int
	for _, m := range a {
		for _, n := range b {
			if m == n {
				matching++
			}
		}
	}

	return matching == len(a)
}

func sub(a, b []int) bool {
	if len(a) == 0 {
		return true
	}

	return false
}

func sup(a, b []int) bool {
	if len(a) == 0 {
		return true
	}

	return false
}
