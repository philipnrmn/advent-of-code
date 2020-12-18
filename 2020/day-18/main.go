package main

import . "../lib"

func part1(input []string) (int, error) {
	result := 0
	for _, line := range input {
		calculate(line)
	}
	return result, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func main() {
	Solve(18, part1, part2)
}
