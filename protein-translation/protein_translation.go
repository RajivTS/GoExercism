package protein

import (
	"errors"
)

var ErrStop error = errors.New("strand contains only stop nucleotide")
var ErrInvalidBase error = errors.New("strand contains invalid nucleotide base")
var codonToProteinMap map[string]string = map[string]string{
	"AUG": "Methionine", "UUU": "Phenylalanine", "UUC": "Phenylalanine", "UUA": "Leucine",
	"UUG": "Leucine", "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine", "UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan", "UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	rnaLst := []rune(rna)
	var protienLst []string
	for idx := 0; idx < len(rnaLst)-2; idx += 3 {
		codon := string(rnaLst[idx]) + string(rnaLst[idx+1]) + string(rnaLst[idx+2])
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return nil, err
		} else if err == ErrStop {
			if protienLst == nil {
				return nil, err
			} else {
				return protienLst, nil
			}
		} else {
			protienLst = append(protienLst, protein)
		}
	}
	return protienLst, nil
}

func FromCodon(codon string) (string, error) {
	protein, exists := codonToProteinMap[codon]
	if !exists {
		return "", ErrInvalidBase
	} else if protein == "STOP" {
		return "", ErrStop
	} else {
		return protein, nil
	}
}
