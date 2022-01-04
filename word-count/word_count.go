package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	r, _ := regexp.Compile("[a-z]+'[a-z]+|[a-z]+|[0-9]+")
	matches := r.FindAllString(strings.ToLower(phrase), -1)
	wordFreq := make(map[string]int)
	for _, word := range matches {
		wordFreq[word]++
	}
	return wordFreq
}
