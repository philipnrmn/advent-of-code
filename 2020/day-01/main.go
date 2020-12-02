package main

import (
	"sort"

	. "../lib"
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
	sort.Ints(input)
	for i := 0; i < len(input); i++ {
		n := input[i]
		for j := i + 1; j < len(input); j++ {
			m := input[j]
			for k := j + 1; k < len(input); k++ {
				o := input[k]
				if n+m+o > 2020 {
					break
				}
				if n+m+o == 2020 {
					return n * m * o, nil
				}
			}
		}
	}
	return 0, NOPE

}

func main() {
	Solve(1, WithInts(part1), WithInts(part2))
}
