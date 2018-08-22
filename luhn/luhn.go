package luhn

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func rToi(r rune) int {
	buf := make([]byte, 1)
	_ = utf8.EncodeRune(buf, r)
	tmp, _ := strconv.Atoi(string(buf))
	return tmp
}

// Valid it return valid status of data
func Valid(s string) bool {

	s = strings.Replace(s, " ", "", -1)

	if len(s) <= 1 {
		return false
	}

	sum := 0

	rs := []rune(s)

	for i := 0; i < len(rs); i++ {

		r := rs[len(rs)-1-i]
		if unicode.IsNumber(r) {
			v := rToi(r)
			if i%2 == 1 {
				v *= 2
				if v > 9 {
					v -= 9
				}
			}
			sum += v
		} else {
			sum = -1
			break
		}
	}

	if sum == -1 {
		return false
	}

	return sum%10 == 0

}
