package main

import "fmt"

func result(numRows int) [][]int {
	segitiga := make([][]int, numRows)
	for i := range segitiga {
		segitiga[i] = make([]int, i+1)
		segitiga[i][0] = 1
		segitiga[i][i] = 1
		for j := 1; j < i; j++ {
			segitiga[i][j] = segitiga[i-1][j-1] + segitiga[i-1][j]
		}
	}
	return segitiga
}

func main() {
	numRows := 5
	segitiga := result(numRows)
	fmt.Println(segitiga)
}
