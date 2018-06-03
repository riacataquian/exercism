// package grains calculate the number of grains of wheat on a chessboard given that the number on each square doubles.
package grains

import (
	"errors"
)

var squares map[int]uint64
var total uint64

func init() {
	calculateSquares()
}

// Square returns the count of grains given a chessboard number.
func Square(n int) (uint64, error) {
	// Zero, negative numbers and numbers greater than 64 are considered invalid numbers.
	if n == 0 || n < 0 || n > 64 {
		return 0, errors.New("invalid number")
	}

	return squares[n], nil
}

// Total returns the total number of grains.
func Total() uint64 {
	return total
}

func calculateSquares() map[int]uint64 {
	n := uint64(1)
	squares = make(map[int]uint64)

	// initial total value
	total = n

	for i := 1; i <= 64; i++ {
		squares[i] = n
		n += n
		total += n
	}

	return squares
}
