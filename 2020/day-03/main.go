package main

import . "../lib"

func part1(input []string) (int, error) {
	var count int

	width := len(input[0])
	x := 0
	for y := 0; y < len(input); y++ {
		r := input[y][x%width]
		if r == '#' {
			count++
		}

		x += 3
	}
	return count, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func main() {
	Solve(3, part1, part2)
}
