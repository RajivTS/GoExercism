package scrabble

import "strings"

func GetLetterScores() map[rune]int {
	scoreMap := make(map[rune]int)
	scoreMap['A'] = 1
	scoreMap['E'] = 1
	scoreMap['I'] = 1
	scoreMap['O'] = 1
	scoreMap['U'] = 1
	scoreMap['L'] = 1
	scoreMap['N'] = 1
	scoreMap['R'] = 1
	scoreMap['S'] = 1
	scoreMap['T'] = 1
	scoreMap['D'] = 2
	scoreMap['G'] = 2
	scoreMap['B'] = 3
	scoreMap['C'] = 3
	scoreMap['M'] = 3
	scoreMap['P'] = 3
	scoreMap['F'] = 4
	scoreMap['H'] = 4
	scoreMap['V'] = 4
	scoreMap['W'] = 4
	scoreMap['Y'] = 4
	scoreMap['K'] = 5
	scoreMap['J'] = 8
	scoreMap['X'] = 8
	scoreMap['Q'] = 10
	scoreMap['Z'] = 10
	return scoreMap
}

func Score(word string) int {
	scoreMap := GetLetterScores()
	var finalScore int
	for _, letter := range strings.ToUpper(word) {
		finalScore += scoreMap[letter]
	}
	return finalScore
}
