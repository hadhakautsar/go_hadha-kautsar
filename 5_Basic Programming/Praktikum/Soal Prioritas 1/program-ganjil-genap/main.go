package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Angka Genap")
		} else {
			fmt.Println(i, "Angka Ganjil")
		}
	}
}
