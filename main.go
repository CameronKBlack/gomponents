package main

import (
	"fmt"

	"gomponents/maths"
)

func main() {
	primes := maths.FindXPrimes(10)
	fmt.Println(primes)
}