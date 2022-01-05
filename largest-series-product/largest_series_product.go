package lsproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	currWinSize, digitLst, errInvalidInput := 0, []rune(digits), errors.New("invalid input")
	if span < 0 || span > len(digits) {
		return 0, errInvalidInput
	}
	if len(digits) == 0 {
		return 1, nil
	}
	var currProduct, maxProduct int64
	for idx, digit := range digitLst {
		if !unicode.IsDigit(digit) {
			return 0, errInvalidInput
		}
		if currProduct == 0 {
			currProduct = int64(digit - '0')
		} else {
			currProduct *= int64(digit - '0')
		}
		currWinSize++
		if digit == '0' {
			currWinSize = 0
		} else if currWinSize >= span {
			if currWinSize > span {
				currProduct /= int64(digitLst[idx-span] - '0')
			}
			if currProduct > maxProduct {
				maxProduct = currProduct
			}
		}
	}
	return maxProduct, nil
}
