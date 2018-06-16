// Package tournament ...
package tournament

import (
	"bytes"
	"fmt"
	"io"
)

// Tally ...
func Tally(r io.Reader, w io.Writer) error {
	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return err
	}

	fmt.Println(b.String())

	return nil
}
