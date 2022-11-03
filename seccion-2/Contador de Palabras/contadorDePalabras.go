package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

func main() {
	i := WordCount("I ate a donut. Then I ate another donut.")
	fmt.Println(i)
}
