package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Println(i)
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go printMultiples(3)
	time.Sleep(90 * time.Second)
}
