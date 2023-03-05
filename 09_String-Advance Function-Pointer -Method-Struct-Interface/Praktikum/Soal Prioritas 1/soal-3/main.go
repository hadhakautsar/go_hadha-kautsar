package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var result string
	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			substring := a[i:j]
			if strings.Contains(b, substring) && len(substring) > len(result) {
				result = substring
			}
		}
	}
	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}
