package main

import "fmt"

func main() {
	//Dengan Perulangan
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Angka Genap")
		} else {
			fmt.Println(i, "Angka Ganjil")
		}
	}

	//Dengan Percabangan dan Scanln untuk input angka
	/* var angka int
	fmt.Println("Masukan Angka: ")
	fmt.Scanln(&angka)

	if angka%2 == 0 {
		fmt.Println(angka, "Angka Genap")
	} else {
		fmt.Println(angka, "Angka Ganjil")
	} */
}
