package main

import (
	. "../lib"
)

func part1(input []string) (int, error) {
	var sum int
	chunks := Chunk(input)

	for _, chunk := range chunks {
		m := map[rune]int{}
		for _, line := range chunk {
			for _, r := range line {
				m[r] = 1
			}
		}
		sum += len(m)
	}

	return sum, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func main() {
	Solve(06, part1, part2)
}
