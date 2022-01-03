package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	tranformedMap := make(map[string]int)
	for score, letters := range in {
		for _, letter := range letters {
			tranformedMap[strings.ToLower(letter)] = score
		}
	}
	return tranformedMap
}
