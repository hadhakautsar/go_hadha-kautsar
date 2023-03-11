package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	fmt.Println("Program cetak bilangan kelipatan 3 yang menerapkan mutex")
	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if num%3 == 0 {
				mu.Lock()
				defer mu.Unlock()
				fmt.Println(num)
			}
		}(i)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
