package engine

import (
	"bufio"
	"bytes"
	"net/http"
	"regexp"
	"strings"
)

// ScanHits is the type for words and their score
type ScanHits = map[string]int

func countWord(text, word string) (count int) {
	scanner := bufio.NewScanner(bytes.NewBufferString(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text() == word {
			count++
		}
	}

	return
}

// Scan returns an array of words
func Scan(text string) ScanHits {
	words := make(map[string]int)

	scanner := bufio.NewScanner(bytes.NewBufferString(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		re, _ := regexp.Compile("[,|.|(|)|_]")
		transformedWord := re.ReplaceAllString(strings.ToLower(word), "")
		words[transformedWord] = countWord(text, word)
	}

	return words
}

// IsText returns whever a file is a
// reconized text file or not.
func IsText(file []byte) bool {
	contentType := http.DetectContentType(file)

	return strings.Contains(contentType, "text/plain")
}

// GetFirstMatchingLine returns the first line to match the given word.
func GetFirstMatchingLine(text string, word string) string {
	scanner := bufio.NewScanner(bytes.NewBufferString(text))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), word) {
			return scanner.Text()
		}
	}

	return ""
}
