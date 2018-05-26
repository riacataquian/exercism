// package diffsquares calculates for the sum of squares, square of sums and their difference.
package diffsquares

// SquareOfSums returns the square of the sum of n.
//
// Given 5, (5 + 4 + ...)^2 = N^2
func SquareOfSums(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i
	}

	return sum * sum
}

// SumOfSquares returns the sum of square.
//
// Given 5, 5^2 + 4^2 + ... = n.
func SumOfSquares(n int) int {
	var sum int

	for i := n; i > 0; i-- {
		sum += i * i
	}

	return sum
}

// Difference returns the difference between a number's square of sums and sums of squares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
