package engine

import "strings"

// CountWord returns the number of
// occurence for a word in a given text.
func CountWord(text, word string) int {
	count := 0

	for _, v := range strings.Split(text, " ") {
		if v == word {
			count++
		}
	}

	return count
}
