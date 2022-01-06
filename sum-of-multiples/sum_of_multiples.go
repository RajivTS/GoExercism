package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum, seen := 0, make(map[int]bool)
	for idx := 0; idx < len(divisors); idx++ {
		if divisors[idx] == 0 {
			continue
		}
		for multiple := 1; multiple*divisors[idx] < limit; multiple++ {
			product := multiple * divisors[idx]
			if !seen[product] {
				sum += product
				seen[product] = true
			}
		}
	}
	return sum
}
