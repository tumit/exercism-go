package scrabble

import "strings"

// Score it return score of scrabble
func Score(input string) int {

	sum := 0

	for _, c := range input {
		sum = sum + aScore(strings.ToUpper(string(c)))
	}

	return sum
}

type strScore struct {
	str   string
	score int
}

var strScores = []strScore{
	{"AEIOULNRST", 1},
	{"DG", 2},
	{"BCMP", 3},
	{"FHVWY", 4},
	{"K", 5},
	{"JX", 8},
	{"QZ", 10},
}

func aScore(s string) int {

	for _, strScore := range strScores {
		if strings.Contains(strScore.str, s) {
			return strScore.score
		}
	}

	return 0
}
