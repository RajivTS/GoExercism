package anagram

import "strings"

func CreateCodec(source string) []int {
	codec := make([]int, 26)
	for _, element := range strings.ToUpper(source) {
		if element >= 'A' && element <= 'Z' {
			idx := int(element - 'A')
			codec[idx]++
		}
	}
	return codec
}

func IsAMatch(source, target []int) bool {
	for idx := 0; idx < len(source); idx++ {
		if source[idx] != target[idx] {
			return false
		}
	}
	return true
}

func Detect(subject string, candidates []string) []string {
	sourceCodec := CreateCodec(subject)
	var matches []string
	for _, target := range candidates {
		targetCodec := CreateCodec(target)
		if IsAMatch(sourceCodec, targetCodec) && !strings.EqualFold(subject, target) {
			matches = append(matches, target)
		}
	}
	return matches
}
