package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var encoded []rune
	segment := 0
	for _, letter := range strings.ToLower(s) {
		if unicode.IsDigit(letter) || unicode.IsLetter(letter) {
			if segment > 0 && segment%5 == 0 {
				encoded = append(encoded, ' ')
				segment = 0
			}
			if unicode.IsLetter(letter) {
				letter = 'z' + 'a' - letter
			}
			encoded = append(encoded, letter)
			segment++
		}
	}
	return string(encoded)
}
