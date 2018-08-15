package collatzconjecture

import (
	"fmt"
)

// CollatzConjecture it return steps count
func CollatzConjecture(i int) (int, error) {

	if i <= 0 {
		return 0, fmt.Errorf("zero is an error")
	}

	_, count := collatzConjecture(i, 0)
	return count, nil
}

func collatzConjecture(i, count int) (int, int) {
	if i == 1 {
		return 0, count
	} else if i%2 == 0 {
		return collatzConjecture(i/2, count+1)
	} else {
		return collatzConjecture((i*3)+1, count+1)
	}
}
