package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToUpper(input)
	seenLetters := make([]bool, 26)
	counter := 0
	for _, letter := range input {
		if letter >= 'A' && letter <= 'Z' {
			idx := int(letter - 'A')
			if !seenLetters[idx] {
				seenLetters[idx] = true
				counter++
			}
		}
	}
	return counter == 26
}
