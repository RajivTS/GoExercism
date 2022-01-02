package strand

import "unicode/utf8"

var dnaToRnaMap map[rune]rune = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

func ToRNA(dna string) string {
	rna := make([]rune, utf8.RuneCountInString(dna))
	for idx, nucleo := range dna {
		rna[idx] = dnaToRnaMap[nucleo]
	}
	return string(rna)
}
