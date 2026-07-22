package differenceofsquares

import "math"

// SquareOfSum returns the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	var sum float64
	for i := range n + 1 {
		sum += float64(i)
	}

	return int(math.Pow(sum, 2))
}

// SumOfSquares returns the sum of the square of the first N natural numbers
func SumOfSquares(n int) int {
	var sum float64
	for i := range n + 1 {
		sum += math.Pow(float64(i), 2)
	}

	return int(sum)
}

// Difference returns the diff beetween SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
