package raindrops

import "strconv"

func Convert(number int) string {
	var message string
	if number%3 == 0 {
		message += "Pling"
	}
	if number%5 == 0 {
		message += "Plang"
	}
	if number%7 == 0 {
		message += "Plong"
	}
	if message == "" {
		return strconv.Itoa(number)
	} else {
		return message
	}
}
