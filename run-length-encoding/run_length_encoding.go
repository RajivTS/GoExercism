package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	lastLetter, count, rleStr, letterAndCount := '0', 0, "", ""
	input += "0"
	for _, letter := range input {
		if letter != lastLetter {
			if count > 1 {
				letterAndCount += strconv.Itoa(count)
			}
			if count > 0 {
				letterAndCount += string(lastLetter)
			}
			lastLetter = letter
			count = 1
			rleStr += letterAndCount
			letterAndCount = ""
		} else {
			count++
		}
	}
	return rleStr
}

func RunLengthDecode(input string) string {
	count, rldStr := 0, ""
	for _, element := range input {
		if unicode.IsDigit(element) {
			count = count*10 + int(element-'0')
		} else if count == 0 {
			rldStr += string(element)
		} else {
			rldStr += strings.Repeat(string(element), count)
			count = 0
		}
	}
	return rldStr
}
