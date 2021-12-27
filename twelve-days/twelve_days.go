package twelve

import (
	"fmt"
	"strings"
)

var giftMap map[int]string
var dayMap map[int]string

func Initialize() {
	if giftMap != nil && dayMap != nil {
		return
	}
	giftMap = make(map[int]string)
	giftMap[1] = "a Partridge in a Pear Tree."
	giftMap[2] = fmt.Sprintf("two Turtle Doves, and %s", giftMap[1])
	giftMap[3] = fmt.Sprintf("three French Hens, %s", giftMap[2])
	giftMap[4] = fmt.Sprintf("four Calling Birds, %s", giftMap[3])
	giftMap[5] = fmt.Sprintf("five Gold Rings, %s", giftMap[4])
	giftMap[6] = fmt.Sprintf("six Geese-a-Laying, %s", giftMap[5])
	giftMap[7] = fmt.Sprintf("seven Swans-a-Swimming, %s", giftMap[6])
	giftMap[8] = fmt.Sprintf("eight Maids-a-Milking, %s", giftMap[7])
	giftMap[9] = fmt.Sprintf("nine Ladies Dancing, %s", giftMap[8])
	giftMap[10] = fmt.Sprintf("ten Lords-a-Leaping, %s", giftMap[9])
	giftMap[11] = fmt.Sprintf("eleven Pipers Piping, %s", giftMap[10])
	giftMap[12] = fmt.Sprintf("twelve Drummers Drumming, %s", giftMap[11])

	dayMap = make(map[int]string)
	dayMap[1] = "first"
	dayMap[2] = "second"
	dayMap[3] = "third"
	dayMap[4] = "fourth"
	dayMap[5] = "fifth"
	dayMap[6] = "sixth"
	dayMap[7] = "seventh"
	dayMap[8] = "eighth"
	dayMap[9] = "ninth"
	dayMap[10] = "tenth"
	dayMap[11] = "eleventh"
	dayMap[12] = "twelfth"
}
func Verse(i int) string {
	Initialize()
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", dayMap[i], giftMap[i])
}

func Song() string {
	verseLst := make([]string, 0)
	for i := 1; i <= 12; i++ {
		verseLst = append(verseLst, Verse(i))
	}
	return strings.Join(verseLst, "\n")
}
