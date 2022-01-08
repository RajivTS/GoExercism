package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	letters := []rune(plain)
	for idx, letter := range letters {
		if unicode.IsLetter(letter) {
			base := 'a'
			if unicode.IsUpper(letter) {
				base = 'A'
			}
			offset := (int(letter-base) + shiftKey) % 26
			letters[idx] = base + rune(offset)
		}
	}
	return string(letters)
}
