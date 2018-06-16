// Package tournament ...
package tournament

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Tally ...
func Tally(r io.Reader, w io.Writer) error {
	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return err
	}

	strs := strings.Split(b.String(), "\n")
	var matches [][]string
	for _, match := range strs {
		m := strings.Split(match, ";")
		if len(m) == 1 && m[0] == "" {
			continue
		}
		matches = append(matches, m)
	}
	fmt.Println(matches)

	return nil
}
