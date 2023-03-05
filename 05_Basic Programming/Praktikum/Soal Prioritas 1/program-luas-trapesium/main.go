package main

import "fmt"

func main() {

	//Program Luas Trapersium
	alasA := 12
	alasB := 16
	t := 5
	luas := (alasA + alasB) * t / 2

	fmt.Println("Luas Trapesium adalah", luas)

	//Menggunakan Scanln untuk input
	/* var alasA, alasB, t int

	fmt.Println("Masukan alasA: ")
	fmt.Scanln(&alasA)
	fmt.Println("Masukan alasB: ")
	fmt.Scanln(&alasB)
	fmt.Println("Masukan Tinggi: ")
	fmt.Scanln(&t)

	luas := (alasA + alasB) * t / 2
	fmt.Println("Luas Trapesium adalah:", luas) */
}
