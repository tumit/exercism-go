package isogram

import (
	"unicode"
)

// IsIsogram it return is isogram
func IsIsogram(s string) bool {

	foundMap := make(map[rune]bool)

	for _, r := range s {

		if !unicode.IsLetter(r) {
			continue
		}

		if _, found := foundMap[unicode.ToUpper(r)]; found {
			return false
		}

		foundMap[unicode.ToUpper(r)] = true
	}

	return true
}
