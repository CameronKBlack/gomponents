package maths

import "math"

func FindXPrimes(x int) []int {
	var primes []int
	n := 2
	for len(primes) < x {
		isPrime := true
		for i := 2; i < int(math.Pow(float64(n), 0.5))+1; i++ {
			if (n % i == 0) {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, n)
		}
		n++
	}
	return primes
}