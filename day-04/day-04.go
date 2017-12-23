package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
}

func solve1(input string) int {
	passphrases := strings.Split(input, "\n")
	count := 0
	for _, pp := range passphrases {
		if isValid1(pp) {
			count++
		}
	}
	return count
}

func isValid1(s string) bool {
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

func solve2(input string) int {
	passphrases := strings.Split(input, "\n")
	count := 0
	for _, pp := range passphrases {
		if isValid2(pp) {
			count++
		}
	}
	return count
}

func isValid2(s string) bool {
	words := strings.Split(s, " ")
	wmap := map[string]bool{}
	for _, w := range words {
		chars := strings.Split(w, "")
		sort.Strings(chars)
		abcw := strings.Join(chars, "")
		if _, found := wmap[abcw]; found {
			return false
		}
		wmap[abcw] = true
	}
	return true
}
