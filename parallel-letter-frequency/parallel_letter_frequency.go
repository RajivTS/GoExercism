package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func FrequencyConc(s string, freqCh chan<- FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	freqCh <- m
}

func MergeMaps(mapA FreqMap, mapB FreqMap) FreqMap {
	for key := range mapB {
		mapA[key] += mapB[key]
	}
	return mapA
}

func CombineResult(chanList []chan FreqMap, lo int, hi int) FreqMap {
	if lo >= hi {
		return <-chanList[lo]
	}
	mid := lo + (hi-lo)/2
	mapLeft := CombineResult(chanList, lo, mid)
	mapRight := CombineResult(chanList, mid+1, hi)
	return MergeMaps(mapLeft, mapRight)
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	chanList := make([]chan FreqMap, len(l))
	for idx, line := range l {
		chanList[idx] = make(chan FreqMap)
		go FrequencyConc(line, chanList[idx])
	}
	return CombineResult(chanList, 0, len(l)-1)
}
