package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat n: ")
	fmt.Scanln(&n)

	ans := make([]string, n+1)
	ans[0] = "0"

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			ans[i] = ans[i-1] + "1"
		} else {
			ans[i] = ans[i/2] + "0"
		}
	}

	fmt.Println(ans)
}
