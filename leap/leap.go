// Package leap contains logic to determine if a given year is leap or not
package leap

// Implementation responsible for determining if a year is leap
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 == 0 {
		return year%400 == 0
	} else {
		return true
	}

}
