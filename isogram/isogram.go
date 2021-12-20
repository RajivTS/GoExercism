package isogram

import "strings"

func IsIsogram(word string) bool {
	seenLetters := make(map[rune]bool)
	for _, letter := range strings.ToLower(word) {
		if seenLetters[letter] && letter != ' ' && letter != '-' {
			return false
		} else {
			seenLetters[letter] = true
		}
	}
	return true
}
