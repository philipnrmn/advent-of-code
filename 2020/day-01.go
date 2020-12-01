package main

import (
	"sort"

	. "./lib"
)

func part1(input []int) (int, error) {
	sort.Ints(input)
	for i := 0; i < len(input); i++ {
		n := input[i]
		for j := i + 1; j < len(input); j++ {
			m := input[j]
			if n+m > 2020 {
				break
			}
			if n+m == 2020 {
				return n * m, nil
			}
		}
	}
	return 0, NOPE
}

func part2(input []int) (int, error) {
	return 0, NYI
}

func main() {
	Solve(1, part1, part2)
}
