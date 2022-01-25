package cipher

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
type shift struct {
	distance int
}

type vigenere struct {
	key []rune
}

// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance < 0 {
		distance = 26 + distance
	}
	if distance <= 0 || distance >= 26 {
		return nil
	}
	return shift{distance: distance}
}

func shiftLetter(distance int, letter rune) rune {
	pos := int(letter - 'a')
	pos = (pos + distance) % 26
	return rune(int('a') + pos)
}

func (c shift) Encode(input string) string {
	var encoded []rune
	for _, elem := range strings.ToLower(input) {
		if unicode.IsLetter(elem) {
			encoded = append(encoded, shiftLetter(c.distance, elem))
		}
	}
	return string(encoded)
}

func (c shift) Decode(input string) string {
	c.distance = 26 - c.distance
	result := c.Encode(input)
	return result
}

func NewVigenere(key string) Cipher {
	totalDist := 0
	for _, elem := range key {
		dist := int(elem - 'a')
		if dist < 0 || dist > 26 {
			return nil
		} else {
			totalDist += dist % 26
		}
	}
	if totalDist == 0 {
		return nil
	} else {
		return vigenere{key: []rune(key)}
	}
}

func (v vigenere) Encode(input string) string {
	var encoded []rune
	keyIdx := 0
	for _, elem := range strings.ToLower(input) {
		if unicode.IsLetter(elem) {
			distance := int(v.key[keyIdx] - 'a')
			encoded = append(encoded, shiftLetter(distance, elem))
			keyIdx = (keyIdx + 1) % len(v.key)
		}
	}
	return string(encoded)
}

func (v vigenere) Decode(input string) string {
	var decoded []rune
	keyIdx := 0
	for _, elem := range strings.ToLower(input) {
		if unicode.IsLetter(elem) {
			distance := 26 - int(v.key[keyIdx]-'a')
			decoded = append(decoded, shiftLetter(distance, elem))
			keyIdx = (keyIdx + 1) % len(v.key)
		}
	}
	return string(decoded)
}
