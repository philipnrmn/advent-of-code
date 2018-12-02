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
	doubles := 0
	triples := 0
	for _, id := range strings.Fields(input) {
		cset := map[rune]int{}
		for _, c := range id {
			cset[c]++
		}
		for _, n := range cset {
			if n == 2 {
				doubles++
				break
			}
		}
		for _, n := range cset {
			if n == 3 {
				triples++
				break
			}
		}
	}

	return doubles * triples
}
