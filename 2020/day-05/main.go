package main

import (
	. "../lib"
)

func part1(input []string) (int, error) {
	highest := 0
	for _, bp := range input {
		instructions := []rune(bp)
		row := partition(instructions[0:7], 0, 127)
		col := partition(instructions[7:], 0, 7)
		id := row*8 + col
		if id > highest {
			highest = id
		}
	}
	return highest, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func partition(instructions []rune, lower, upper int) int {
	if len(instructions) == 0 {
		return lower
	}
	i := instructions[0]
	delta := upper - lower
	if i == 'L' || i == 'F' {
		return partition(instructions[1:], lower, lower+(delta/2))
	}
	return partition(instructions[1:], upper-(delta/2), upper)
}

func main() {
	Solve(5, part1, part2)
}
