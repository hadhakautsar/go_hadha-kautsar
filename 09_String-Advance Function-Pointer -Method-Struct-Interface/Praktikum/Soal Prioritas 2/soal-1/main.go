package main

import "fmt"

func caesar(offset int, input string) string {
	var output string

	for _, char := range input {
		if char == ' ' {
			output += " "
			continue
		}

		geser := int(char) - 'a' + offset
		geser = geser%26 + 'a'
		output += string(geser)
	}

	return output
}

func main() {
	fmt.Println(caesar(3, "abc"))                           //def
	fmt.Println(caesar(2, "alta"))                          //cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
