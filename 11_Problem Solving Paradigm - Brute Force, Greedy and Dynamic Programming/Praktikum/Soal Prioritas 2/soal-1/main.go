package main

import "fmt"

func Frog(jumps []int) int {
	n := len(jumps)
	biayaMin := make([]int, n)
	biayaMin[0] = 0
	if jumps[1] > jumps[0] {
		biayaMin[1] = jumps[1] - jumps[0]
	} else {
		biayaMin[1] = jumps[0] - jumps[1]
	}

	for i := 2; i < n; i++ {
		biaya1 := biayaMin[i-1]
		if jumps[i] > jumps[i-1] {
			biaya1 += jumps[i] - jumps[i-1]
		} else {
			biaya1 += jumps[i-1] - jumps[i]
		}
		biaya2 := biayaMin[i-2]
		if jumps[i] > jumps[i-2] {
			biaya2 += jumps[i] - jumps[i-2]
		} else {
			biaya2 += jumps[i-2] - jumps[i]
		}
		if biaya1 < biaya2 {
			biayaMin[i] = biaya1
		} else {
			biayaMin[i] = biaya2
		}
	}

	return biayaMin[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
