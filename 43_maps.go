package main

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) (wordCount map[string]int) {
	words := strings.Fields(s)
	wordCount = make(map[string]int)
	fmt.Println(words)
	for _, word := range words {
		wordCount[word]++
	}

	return
}

func main() {
	wc.Test(WordCount)
}
