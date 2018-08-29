package proverb

import "fmt"

// Proverb it return generated of the relevant proverb
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return []string{}
	}

	proverbs := make([]string, len(rhyme))

	if len(rhyme) > 1 {
		for i := range rhyme[:len(rhyme)-1] {
			proverbs[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
	}

	proverbs[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return proverbs
}
