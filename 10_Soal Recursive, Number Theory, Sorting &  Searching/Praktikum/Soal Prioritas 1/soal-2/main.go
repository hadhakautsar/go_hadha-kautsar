package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) string {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item]++
	}

	pairs := make([]pair, 0, len(counts))
	for k, v := range counts {
		pairs = append(pairs, pair{k, v})
	}
	// Bubble Sort
	for i := 0; i < len(pairs)-1; i++ {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[i].count < pairs[j].count {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}

	var result string
	for _, p := range pairs {
		result += fmt.Sprintf("%s->%d ", p.name, p.count)
	}
	return result
}
func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1 basketball->1 tenis->1
}
