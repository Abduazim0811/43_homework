package wordcount

import "strings"

func Wordcount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}
