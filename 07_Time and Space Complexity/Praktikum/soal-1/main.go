package main

import "fmt"

func primeNumber(number int) bool {
	numPrime := true
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	if numPrime == true {
		return true
	}
	return numPrime
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
