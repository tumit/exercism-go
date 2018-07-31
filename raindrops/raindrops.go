package raindrops

import "strconv"

// Convert return
func Convert(n int) (r string) {

	if n%105 == 0 {
		return "PlingPlangPlong"
	}

	if n%35 == 0 {
		return "PlangPlong"
	}

	if n%21 == 0 {
		return "PlingPlong"
	}

	if n%15 == 0 {
		return "PlingPlang"
	}

	if n%3 == 0 {
		return "Pling"
	}

	if n%5 == 0 {
		return "Plang"
	}

	if n%7 == 0 {
		return "Plong"
	}

	return strconv.Itoa(n)
}
