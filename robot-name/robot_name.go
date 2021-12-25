package robotname

import (
	"errors"
	"math/rand"
	"strconv"
)

// Define the Robot type here.
type Robot struct {
	ID string
}

var alphabetPool []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var nameMap map[string]bool = make(map[string]bool)
var maxNamesCount int = 26 * 26 * 10 * 10 * 10

func (r *Robot) Name() (string, error) {
	errAllNamesTaken := errors.New("all robot names already taken")
	if r.ID != "" {
		return r.ID, nil
	} else if len(nameMap) == maxNamesCount {
		return "", errAllNamesTaken
	}
	for {
		fstLetter := alphabetPool[rand.Intn(26)]
		sndLetter := alphabetPool[rand.Intn(26)]
		fstDigit := rand.Intn(10)
		sndDigit := rand.Intn(10)
		thrdDigit := rand.Intn(10)
		newName := string(fstLetter) + string(sndLetter) + strconv.Itoa(fstDigit) + strconv.Itoa(sndDigit) + strconv.Itoa(thrdDigit)
		if !nameMap[newName] {
			r.ID = newName
			nameMap[newName] = true
			return newName, nil
		}
	}
}

func (r *Robot) Reset() {
	r.ID = ""
}
