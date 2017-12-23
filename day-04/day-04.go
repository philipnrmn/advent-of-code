package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) int {
	passphrases := strings.Split(input, "\n")
	count := 0
	for _, pp := range passphrases {
		if isValid(pp) {
			count++
		}
	}
	return count
}

func isValid(s string) bool {
	words := strings.Split(s, " ")
	wmap := map[string]bool{}
	for _, w := range words {
		if _, found := wmap[w]; found {
			return false
		}
		wmap[w] = true
	}
	return true
}
