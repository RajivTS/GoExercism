package romannumerals

import "errors"

type NumericValue struct {
	Decimal int
	Roman   string
}

func GetNumericValues() []NumericValue {
	values := make([]NumericValue, 13)
	values[0] = NumericValue{Decimal: 1, Roman: "I"}
	values[1] = NumericValue{Decimal: 4, Roman: "IV"}
	values[2] = NumericValue{Decimal: 5, Roman: "V"}
	values[3] = NumericValue{Decimal: 9, Roman: "IX"}
	values[4] = NumericValue{Decimal: 10, Roman: "X"}
	values[5] = NumericValue{Decimal: 40, Roman: "XL"}
	values[6] = NumericValue{Decimal: 50, Roman: "L"}
	values[7] = NumericValue{Decimal: 90, Roman: "XC"}
	values[8] = NumericValue{Decimal: 100, Roman: "C"}
	values[9] = NumericValue{Decimal: 400, Roman: "CD"}
	values[10] = NumericValue{Decimal: 500, Roman: "D"}
	values[11] = NumericValue{Decimal: 900, Roman: "CM"}
	values[12] = NumericValue{Decimal: 1000, Roman: "M"}
	return values
}

func ToRomanNumeral(input int) (string, error) {
	errOutOfRange := errors.New("input number is out of range")
	if input <= 0 || input > 3000 {
		return "", errOutOfRange
	}
	values := GetNumericValues()
	idx := 12
	var result string
	for input > 0 {
		div := input / values[idx].Decimal
		input = input % values[idx].Decimal
		for div > 0 {
			result += values[idx].Roman
			div--
		}
		idx--
	}
	return result, nil
}
