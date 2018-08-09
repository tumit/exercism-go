package isogram

import "strings"

var ignores = "- "

func removeIgnoreString(s string) string {
	for _, ignore := range ignores {
		s = strings.Replace(s, string(ignore), "", -1)
	}
	return s
}

// IsIsogram it return is isogram
func IsIsogram(s string) bool {

	s = strings.ToUpper(removeIgnoreString(s))

	for _, r := range s {
		if strings.Count(s, string(r)) != 1 {
			return false
		}
	}

	return true
}
