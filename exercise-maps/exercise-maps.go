package main

import (
	"strings"

	"golang.org/x/tour/wc"
)
//WordCount counts the words
func WordCount(s string) map[string]int {
	var words = make(map[string]int)
	for _, v := range strings.Fields(s) {
		if words[v] == 0 {
			words[v] = 1
		} else {
			words[v]++
		}
	}
	return words
}

func main() {
	//fmt.Println(strings.Fields("mother washes window"))
	wc.Test(WordCount)
}
