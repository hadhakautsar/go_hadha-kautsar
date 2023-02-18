package main

import "fmt"

func main() {
	var nilai int = 84
	if nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else {
		switch {
		case nilai >= 80:
			fmt.Println("Nilai A")
		case nilai >= 65:
			fmt.Println("Nilai B")
		case nilai >= 50:
			fmt.Println("Nilai C")
		case nilai >= 35:
			fmt.Println("Nilai D")
		case nilai >= 0:
			fmt.Println("Nilai E")
		default:
			fmt.Println("Nilai Invalid")
		}
	}

	//Menggunakan Scanln untuk memasukan input nilai
	/* var nilai int
	fmt.Println("Masukan Angka :")
	fmt.Scanln(&nilai)

	if nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else {
		switch {
		case nilai >= 80:
			fmt.Println("Nilai A")
		case nilai >= 65:
			fmt.Println("Nilai B")
		case nilai >= 50:
			fmt.Println("Nilai C")
		case nilai >= 35:
			fmt.Println("Nilai D")
		case nilai >= 0:
			fmt.Println("Nilai E")
		default:
			fmt.Println("Nilai Invalid")
		}
	} */
}
