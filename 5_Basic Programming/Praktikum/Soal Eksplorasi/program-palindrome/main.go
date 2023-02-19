package main

import "fmt"

func Palindrome(cek string) bool {
	for i := 0; i < len(cek)/2; i++ {
		if cek[i] != cek[len(cek)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var kata string
	fmt.Println("Masukan kata: ")
	fmt.Scanln(&kata)
	output := Palindrome(kata)
	if output == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}
