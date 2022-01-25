package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func GetNormalizedStr(plainText string) []rune {
	var output []rune
	for _, elem := range strings.ToLower(plainText) {
		if unicode.IsLetter(elem) || unicode.IsDigit(elem) {
			output = append(output, elem)
		}
	}
	return output
}

func GetDimensions(textLength int) (int, int) {
	col := int(math.Sqrt(float64(textLength)))
	row := col - 1
	for col*row < textLength {
		if row < col {
			row++
		} else {
			col++
		}
	}
	return col, row
}

func Encode(pt string) string {
	text := GetNormalizedStr(pt)
	col, row := GetDimensions(len(text))
	var encodedMsg []string
	for c := 0; c < col; c++ {
		chunk := ""
		for r := 0; r < row; r++ {
			idx := r*col + c
			if idx >= len(text) {
				chunk += " "
			} else {
				chunk += string(text[idx])
			}
		}
		encodedMsg = append(encodedMsg, chunk)
	}
	return strings.Join(encodedMsg, " ")
}
