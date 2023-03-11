package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 30; i++ {
			if i%3 == 0 {
				ch <- i
			}
		}
		close(ch)
	}()

	for num := range ch {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(num)
	}
}
