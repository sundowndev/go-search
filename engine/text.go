package engine

import (
	"bufio"
	"bytes"
	"net/http"
	"regexp"
	"strings"
)

// CountWord returns the number of
// occurence for a word in a given text.
func CountWord(text, word string) (count int) {
	for _, v := range GetWordsFromText(text) {
		if v == word {
			count++
		}
	}

	return
}

// GetWordsFromText returns an array of words
func GetWordsFromText(text string) (words []string) {
	scanner := bufio.NewScanner(bytes.NewBufferString(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())
		re, _ := regexp.Compile("[,|.|(|)|_]")
		words = append(words, re.ReplaceAllString(w, ""))
	}

	return words
}

// IsTextFile returns whever a file is a
// reconized text file or not.
func IsTextFile(file []byte) bool {
	contentType := http.DetectContentType(file)

	return strings.Index(contentType, "text/plain") > -1
}

// GetFirstMatchingLine returns the first line to match the given word.
func GetFirstMatchingLine(text string, word string) string {
	scanner := bufio.NewScanner(bytes.NewBufferString(text))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.Index(strings.ToLower(scanner.Text()), word) > -1 {
			return scanner.Text()
		}
	}

	return ""
}
