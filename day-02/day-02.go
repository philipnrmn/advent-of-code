package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
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

func solve2(input string) string {
	result := ""
	for j, id1 := range strings.Fields(input) {
		for k, id2 := range strings.Fields(input) {
			if j == k {
				continue
			}
			common := ""
			for i := 0; i < len(id1); i++ {
				if id1[i] == id2[i] {
					common += string(id1[i])
				}
			}
			if len(common) > len(result) {
				result = common
			}
		}
	}
	return result
}
