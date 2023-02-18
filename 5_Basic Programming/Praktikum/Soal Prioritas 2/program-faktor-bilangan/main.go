package main

import "fmt"

func main() {
	var angka int
	fmt.Println("Masukan Angka: ")
	fmt.Scanln(&angka)
	fmt.Println("Faktor Bilangan dari angka", angka, "adalah")

	for i := angka; i > 0; i-- {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}
