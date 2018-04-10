// https://go-tour-ru-ru.appspot.com/moretypes/23

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func main() {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	wordsStat := make(map[string]int)

	for _, word := range words {
		wordsStat[word]++
	}

	return wordsStat
}
