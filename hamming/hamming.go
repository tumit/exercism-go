package hamming

import "fmt"

// Distance return counting how many of the nucleotides
// are different from their equivalent in the other string
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, fmt.Errorf("len(%q), len(%q) should be equal", a, b)
	}

	notMatched := 0

	for i := 0; i < len(a); i++ {
		aChr := a[i]
		bChr := b[i]
		if aChr != bChr {
			notMatched++
		}
	}

	return notMatched, nil
}
