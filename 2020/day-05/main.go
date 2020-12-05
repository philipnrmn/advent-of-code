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
	seats := make([][]int, 128)
	for i := 0; i < 128; i++ {
		seats[i] = make([]int, 8)
	}

	for _, bp := range input {
		instructions := []rune(bp)
		row := partition(instructions[0:7], 0, 127)
		col := partition(instructions[7:], 0, 7)
		seats[row][col] = 1
	}
	// this is how I solved it initially
	// for i, row := range seats {
	//   fmt.Printf("%03d %v\n", i, row)
	// }

	// Actual solution:
	sums := make([]int, 128)
	for i := 0; i < 128; i++ {
		sums[i] = sumOf(seats[i])
	}
	for i := 1; i < 127; i++ {
		if sums[i-1] == 8 && sums[i] == 7 && sums[i+1] == 8 {
			for j := 0; j < 7; j++ {
				if seats[i][j] == 0 {
					return i*8 + j, nil
				}
			}
		}
	}

	return 0, NOPE
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

func sumOf(s []int) int {
	sum := s[0]
	for i := 1; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

func main() {
	Solve(5, part1, part2)
}
