package libs

// IsPrime determines whether n is prime. O(√n)
func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GetPrimes returns prime numbers less than or equal to n. O(nlog(log(n)))
func GetPrimes(n int) (primes []int) {
	isPrimes := make([]bool, n+1)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	isPrimes[0] = false
	isPrimes[1] = false

	for i := 2; i*i <= n; i++ {
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

/*
PrimeFactorize performs prime factorization of an integer n.
Returns a map (key: prime factor, value: exponent of prime factor). O(√n)
*/
func PrimeFactorize(n int) map[int]int {
	ret := make(map[int]int)
	rem := n

	for p := 2; p*p <= n; p++ {
		if rem%p != 0 {
			continue
		}

		e := 0
		for rem%p == 0 {
			e++
			rem /= p
		}

		ret[p] = e
	}

	if rem != 1 {
		ret[rem] = 1
	}

	return ret
}
