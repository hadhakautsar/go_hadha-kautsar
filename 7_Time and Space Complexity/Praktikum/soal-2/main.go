package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		y := pow(x, n/2)
		return y * y
	} else {
		y := pow(x, (n-1)/2)
		return x * y * y
	}
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
