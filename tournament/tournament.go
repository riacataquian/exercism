// Package tournament tally the results of a small football competition.
//
// Semantics:
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
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Match represents players and their score tables.
type Match map[string]Table

// Table describes a score board.
type Table struct {
	// Matches Played.
	MP int
	// Matches Won.
	W int
	// Matches Drawn/Tied.
	D int
	// Matches Lost.
	L int
	// Points.
	P int
}

// MatchTally maps players and their score tables.
type MatchTally struct {
	Player string
	Table
}

// Tally ...
func Tally(r io.Reader, w io.Writer) error {
	// Read values from r.
	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return err
	}

	strs := strings.Split(b.String(), "\n")
	matches := make(Match)
	for _, match := range strs {
		// ; separates players and results.
		m := strings.Split(match, ";")

		// Ignore wrong formats.
		if len(m) == 1 || m[0] == "" {
			continue
		}

		// Tests expects an error if len is 2.
		if len(m) == 2 {
			return errors.New("wrong input format")
		}

		// Player 1 and 2 names.
		p1 := strings.TrimSpace(m[0])
		p2 := strings.TrimSpace(m[1])

		// Player 1 and 2 score tables.
		t1 := matches[p1]
		t2 := matches[p2]

		// Increment matches played.
		t1.MP++
		t2.MP++

		// Increment scores accordingly.
		switch m[2] {
		case "win":
			t1.W++
			t1.P += 3

			t2.L++
		case "loss":
			t1.L++

			t2.W++
			t2.P += 3
		case "draw":
			t1.D++
			t2.D++

			t1.P++
			t2.P++
		default:
			return errors.New("unknown results")
		}

		matches[p1] = t1
		matches[p2] = t2
	}

	err := writeTable(w, sortByPoints(matches))
	if err != nil {
		return err
	}

	return nil
}

// writeTable writes tally to w.
func writeTable(w io.Writer, matches []MatchTally) error {
	// Write table header.
	_, err := io.WriteString(w, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}

	// Write table body.
	for _, e := range matches {
		tpl := "%-31s| %2d | %2d | %2d | %2d | %2d\n"
		s := fmt.Sprintf(tpl, e.Player, e.MP, e.W, e.D, e.L, e.P)
		// Output:
		// Player Name               | 1 | 2 | 3 | 4 | 5

		_, err := io.WriteString(w, s)
		if err != nil {
			return err
		}
	}

	return nil
}

// sortByPoints sorts tally by points.
// For draw points, it alphabetically sorts the players name.
func sortByPoints(matches Match) []MatchTally {
	// Convert matches to slices for stable sorting.
	// Maps in Go is being sorted randomly.
	slice := make([]MatchTally, 0, len(matches))
	for k, v := range matches {
		slice = append(slice, MatchTally{k, v})
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
