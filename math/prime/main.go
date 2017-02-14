package main

import (
	"fmt"
	"math"
)

func main() {
	p := primesList(100)
	fmt.Println(p)
}

func primesList(limit int) []int {
	var primes []int
	for i := 0; i <= limit; i++ {
		prime := true

		if i < 2 {
			continue
		}

		sqrt := int(math.Floor(math.Sqrt(float64(i))))

		for x := 2; x < sqrt; x++ {
			if i%x == 0 {
				prime = false
			}
		}

		if prime == true {
			primes = append(primes, i)
		}
	}

	return primes
}
