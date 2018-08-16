package diffsquares

// SquareOfSums it return square of sums
func SquareOfSums(n int) int {
	sum := (n * (1 + n)) / 2
	return sum * sum
}

// SumOfSquares it return sum of squares
func SumOfSquares(n int) int {
	sum := (n * (n + 1) * ((2 * n) + 1)) / 6
	return int(sum)
}

// Difference it return difference of square of sums and sum of square
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
