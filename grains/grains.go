package grains

import (
	"fmt"
)

// Square it return calculate the number of grains of wheat on a chessboard
func Square(i int) (uint64, error) {

	if i == 0 {
		return 0, fmt.Errorf("square 0 returns an error")
	}

	if i < 0 {
		return 0, fmt.Errorf("negative square returns an error")
	}

	if i > 64 {
		return 0, fmt.Errorf("square greater than 64 returns an error")
	}

	return 1 << uint64(i-1), nil
	// return uint64(math.Pow(2.0, float64(i-1))), nil
}

// Total it return total of grains of wheat on a chessboard
func Total() uint64 {
	v, _ := Square(65)
	return v - 1
}
