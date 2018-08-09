package scrabble

import (
	"unicode"
)

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

func makeScoreMap(strScores []strScore) map[rune]int {
	m := make(map[rune]int)
	for _, strScore := range strScores {
		for _, str := range strScore.str {
			m[unicode.ToLower(str)] = strScore.score
			m[unicode.ToUpper(str)] = strScore.score
		}
	}
	return m
}

var scoreMap = makeScoreMap(strScores)

// Score it return score of scrabble
func Score(input string) int {

	sum := 0

	for _, c := range input {
		sum = sum + scoreMap[c]
	}

	return sum
}
