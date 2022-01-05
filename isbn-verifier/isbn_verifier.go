package isbn

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	multiplier, sum, digit, check := 10, 0, 0, 'X'
	isbn = strings.ReplaceAll(isbn, "-", "")
	for idx, letter := range isbn {
		if letter == check && idx != 9 {
			return false
		} else if letter == check {
			digit = 10
		} else {
			digit = int(letter - '0')
		}
		sum = sum + digit*multiplier
		multiplier--
	}
	return len(isbn) == 10 && sum%11 == 0
}
