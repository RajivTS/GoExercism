package pythagorean

import "sort"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var result []Triplet
	for c := max; c > min; c-- {
		for a, b := min, c-1; a < b; {
			squareSum, target := a*a+b*b, c*c
			if squareSum > target {
				b--
			} else if squareSum < target {
				a++
			} else {
				result = append(result, [3]int{a, b, c})
				a++
				b--
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		} else if result[i][1] != result[j][1] {
			return result[i][1] < result[j][1]
		} else {
			return result[i][2] <= result[j][2]
		}
	})
	return result
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var result []Triplet
	for _, triple := range Range(1, p) {
		if triple[0]+triple[1]+triple[2] == p {
			result = append(result, triple)
		}
	}
	return result
}
