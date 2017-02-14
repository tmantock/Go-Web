package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	fn := strings.ToUpper(flag.Arg(0))
	nm, _ := strconv.Atoi(flag.Arg(1))

	switch fn {
	case "LIST":
		result := primesList(nm)
		fmt.Println(result)
	case "PRIME":
		result := isPrime(nm)
		fmt.Println(result)
	case "FACTORS":
		result := primeFactors(nm)
		fmt.Println(result)
	default:
		fmt.Println("Invalid Command")
	}

	os.Exit(0)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	sqrt := int(math.Floor(math.Sqrt(float64(num))))

	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
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

func primeFactors(num int) []int {
	var factors []int
	isPrime := num

	for i := 2; i <= isPrime; i++ {
		if isPrime%i == 0 {
			isPrime = isPrime / i
			factors = append(factors, i)
		}
	}

	return factors
}
