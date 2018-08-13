package bob

import (
	"strings"
	"unicode"
)

func foundLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}

	upper := strings.ToUpper(remark)

	if strings.HasSuffix(remark, "?") {

		if foundLetter(remark[0:len(remark)-1]) && upper == remark {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}

	if foundLetter(remark) && upper == remark {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
