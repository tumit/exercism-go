package isogram

import (
	"strings"
)

var ignores = "- "

func isNotIgnore(r rune) bool {
	return strings.IndexRune(ignores, r) < 0
}

func isRepeat(s string, r rune) bool {
	firstIndex := strings.IndexRune(s, r)
	return strings.IndexRune(s[firstIndex+1:len(s)], r) > -1
}

// IsIsogram it return is isogram
func IsIsogram(s string) bool {

	s = strings.ToUpper(s)

	for _, r := range s {
		if isNotIgnore(r) && isRepeat(s, r) {
			return false
		}
	}

	return true
}
