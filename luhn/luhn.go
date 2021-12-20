package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	chars := []rune(id)
	if len(chars) <= 1 {
		return false
	}
	runningSum := 0
	toDouble := false
	for i := len(chars) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(chars[i]))
		if err != nil {
			return false
		} else if toDouble {
			num = num * 2
			if num > 9 {
				num -= 9
			}
			toDouble = false
		} else {
			toDouble = true
		}
		runningSum += num
	}
	return runningSum%10 == 0
}
