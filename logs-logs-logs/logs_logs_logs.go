package logs

import (
	"math"
	"strings"
	"unicode/utf8"
)

func GetIndexOf(str string, pattern rune) int {
	idx := strings.IndexRune(str, pattern)
	if idx == -1 {
		idx = math.MaxInt
	}
	return idx
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	searchIdx := GetIndexOf(log, '🔍')
	recommendationIdx := GetIndexOf(log, '❗')
	weatherIdx := GetIndexOf(log, '☀')
	if searchIdx < recommendationIdx && searchIdx < weatherIdx {
		return "search"
	} else if recommendationIdx < searchIdx && recommendationIdx < weatherIdx {
		return "recommendation"
	} else if weatherIdx < searchIdx && weatherIdx < recommendationIdx {
		return "weather"
	} else {
		return "default"
	}
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
