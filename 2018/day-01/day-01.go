package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
}

func solve1(input string) int {
	result := 0
	for _, field := range strings.Fields(input) {
		m := 1
		if field[0] == byte('-') {
			m = -1
		}
		n, _ := strconv.Atoi(field[1:])
		result += n * m

	}
	return result
}

func solve2(input string) int {
	fields := strings.Fields(input)
	frequencies := []int{0}
	f := 0
	i := 0
	for {
		field := fields[i%len(fields)]
		m := 1
		if field[0] == byte('-') {
			m = -1
		}
		n, _ := strconv.Atoi(field[1:])
		f += n * m
		if contains(frequencies, f) {
			return f
		}
		frequencies = append(frequencies, f)
		i++
	}
}

func contains(haystack []int, needle int) bool {
	for _, h := range haystack {
		if needle == h {
			return true
		}
	}
	return false
}
