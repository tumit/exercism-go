package raindrops

import "strconv"

// Convert a number to a string, the contents of which depend on the number's factors.
// - If the number has 3 as a factor, output 'Pling'.
// - If the number has 5 as a factor, output 'Plang'.
// - If the number has 7 as a factor, output 'Plong'.
// - If the number does not have 3, 5, or 7 as a factor,
//   just pass the number's digits straight through.
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
