package prime

var primes []int = []int{2, 3, 5, 7, 11, 13}

func IsPrime(num int) bool {
	for _, divisor := range primes {
		if num%divisor == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	} else {
		lastPrimeIdx := len(primes)
		for val := primes[lastPrimeIdx-1]; lastPrimeIdx < n; val++ {
			if IsPrime(val) {
				primes = append(primes, val)
				lastPrimeIdx++
			}
		}
	}
	return primes[n-1], true
}
