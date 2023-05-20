package lib

import "math"

// IsPrime O(√n)
func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GetPrimes O(nloglogn)
func GetPrimes(n int) (primes []int) {
	isPrimes := make([]bool, n+1)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	isPrimes[0] = false
	isPrimes[1] = false

	nSqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= nSqrt; i++ {
		if isPrimes[i] {
			for j := 2 * i; j <= n; j += i {
				isPrimes[j] = false
			}
		}
	}

	for i := range isPrimes {
		if isPrimes[i] {
			primes = append(primes, i)
		}
	}

	return
}
