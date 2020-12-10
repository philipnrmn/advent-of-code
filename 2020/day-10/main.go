package main

import (
	"sort"

	. "../lib"
)

func part1(raw []string) (int, error) {
	input := Ints(raw)
	sort.Ints(input)
	var prev, ones, threes int
	for i := 0; i < len(input); i++ {
		if input[i]-1 == prev {
			ones++
		}
		if input[i]-3 == prev {
			threes++
		}
		prev = input[i]
	}
	return ones * (threes + 1), nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func main() {
	Solve(10, part1, part2)
}
