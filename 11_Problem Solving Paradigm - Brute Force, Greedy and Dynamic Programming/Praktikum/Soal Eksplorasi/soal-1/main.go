package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukkan angka normal: ")
	fmt.Scanln(&num)

	romanNums := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var konversi string
	for i := 0; i < len(romanNums); i++ {
		for num >= romanValues[i] {
			konversi += romanNums[i]
			num -= romanValues[i]
		}
	}
	fmt.Println("Angka Romawi:", konversi)
}
