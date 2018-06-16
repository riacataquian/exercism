// Package tournament ...
//
// Bilthering Badgers win:
// Allegoric Alaskans;Blithering Badgers;win
//
// Courageous Californians win:
// Courageous Californians;Blithering Badgers;loss
//
// Tied:
// Devastating Donkeys;Courageous Californians;draw
//
// A win earns a team 3 points. A draw earns 1. A loss earns 0.
package tournament

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Match ...
type Match map[string]Table

// Table ...
type Table struct {
	// Matches Played
	MP int
	// Matches Won
	W int
	// Matches Drawn/Tied
	D int
	// Matches Lost
	L int
	// Points
	P int
}

// Entry ...
type Entry struct {
	player string
	table  Table
}

// Tally ...
func Tally(r io.Reader, w io.Writer) error {
	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return err
	}

	strs := strings.Split(b.String(), "\n")
	var matches = make(Match)
	for _, match := range strs {
		m := strings.Split(match, ";")
		if len(m) == 1 && m[0] == "" {
			continue
		}

		// Player 1 and 2 names.
		p1 := m[0]
		p2 := m[1]

		// Player 1 and 2 score tables.
		t1 := matches[p1]
		t2 := matches[p2]

		// Increment matches played.
		t1.MP++
		t2.MP++

		switch m[2] {
		case "win":
			t1.W++
			t1.P += 3

			t2.L++
		case "loss":
			t1.L++

			t2.W++
			t2.P += 3
		default: // draw
			t1.D++
			t2.D++

			t1.P++
			t2.P++
		}

		matches[p1] = t1
		matches[p2] = t2
	}

	printSorted(matches)

	return nil
}

// type byPoints []Entry

// func (d byPoints) Len() int {
// 	return len(d)
// }
// func (d byPoints) Less(i, j int) bool {
// 	return d[i].value.Priority < d[j].value.Priority
// }
// func (d byPoints) Swap(i, j int) {
// 	d[i], d[j] = d[j], d[i]
// }

func printSorted(matches Match) {
	slice := make([]Entry, 0, len(matches))
	for k, v := range matches {
		fmt.Println(k)
		slice = append(slice, Entry{k, v})
	}
	fmt.Println(slice)
}
