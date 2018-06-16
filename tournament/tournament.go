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
	"sort"
	"strings"
	// "unicode"
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
	Player string
	Table
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
		if len(m) == 1 || m[0] == "" {
			continue
		}

		// Player 1 and 2 names.
		p1 := strings.TrimSpace(m[0])
		p2 := strings.TrimSpace(m[1])
		// p1 := strings.TrimRightFunc(m[0], func(r rune) bool {
		// 	return unicode.IsSpace(r)
		// })
		// p2 := strings.TrimRightFunc(m[1], func(r rune) bool {
		// 	return unicode.IsSpace(r)
		// })

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

	err := writeTable(w, matches)
	if err != nil {
		return err
	}

	return nil
}

func writeTable(w io.Writer, matches map[string]Table) error {
	_, err := io.WriteString(w, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}

	for _, e := range sortByPoints(matches) {
		tpl := "%-31s| %2d | %2d | %2d | %2d | %2d\n"
		s := fmt.Sprintf(tpl, e.Player, e.MP, e.W, e.D, e.L, e.P)
		_, err := io.WriteString(w, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func sortByPoints(matches Match) []Entry {
	slice := make([]Entry, 0, len(matches))
	for k, v := range matches {
		slice = append(slice, Entry{k, v})
	}
	sort.Slice(slice, func(i, j int) bool {
		prev := slice[i].Table.P
		next := slice[j].Table.P

		// If points are equal, sort player names alphabetically.
		if prev == next {
			return slice[j].Player > slice[i].Player
		}

		return prev > next
	})
	return slice
}
