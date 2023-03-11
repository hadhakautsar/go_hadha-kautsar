package main

import (
	"fmt"
	"sync"
)

func countFrequency(text string, freqMap *map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, char := range text {
		(*freqMap)[char]++
	}
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	freqMap := make(map[rune]int)
	var wg sync.WaitGroup

	splitText := make([]string, 4)
	splitSize := len(text) / 4

	for i := 0; i < 4; i++ {
		startIndex := i * splitSize
		endIndex := startIndex + splitSize

		if i == 3 {
			splitText[i] = text[startIndex:]
		} else {
			splitText[i] = text[startIndex:endIndex]
		}

		wg.Add(1)
		go countFrequency(splitText[i], &freqMap, &wg)
	}

	wg.Wait()

	for char, freq := range freqMap {
		fmt.Printf("%c: %d\n", char, freq)
	}
}
