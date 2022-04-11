package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sList := strings.Split(s, " ")
	sMap := make(map[string]int)

	for _, v := range sList {
		sMap[v]++
	}

	return sMap
}

func main() {
	wc.Test(WordCount)
}
