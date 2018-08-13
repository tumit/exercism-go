package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	// words := strings.Fields(s)
	words := regexp.MustCompile("[ -]").Split(s, -1)

	abbr := make([]string, len(words))

	for i, word := range words {
		abbr[i] = strings.ToUpper(word[0:1])
	}

	return strings.Join(abbr, "")
}
