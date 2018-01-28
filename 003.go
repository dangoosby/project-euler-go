package main

import (
	"fmt"
	"math"
)

func PrimeCheck(val int) bool {
	for i := 2; i <= int(math.Floor(float64(val)/2)); i++ {
		if val%i == 0 {
			return false
		}
	}
	return val > 1
}

func main() {
	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143 ?

	var primes []int
	num := 600851475143
	sr := math.Sqrt(float64(num))
	fmt.Println("Number: ", num)

	for i := 3; i < int(sr); i += 2 {
		if num%i == 0 {
			if PrimeCheck(i) {
				primes = append(primes, i)
			}
		}
	}

	fmt.Println("Largest Prime Factor: ", primes[len(primes)-1])
}
