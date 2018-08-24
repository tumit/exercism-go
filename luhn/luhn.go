package luhn

import (
	"strings"
)

// Valid it return valid status of data
func Valid(s string) bool {

	s = strings.Replace(s, " ", "", -1)

	if len(s) <= 1 {
		return false
	}

	sum := 0
	mod := len(s) % 2

	for i, r := range s {

		v, b := r2i(r)

		if !b {
			return false
		}

		if i%2 == mod {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}

		sum += v
	}

	return sum%10 == 0
}

func r2i(r rune) (int, bool) {
	n := int(r - '0')
	return n, 0 <= n && n <= 9
}
