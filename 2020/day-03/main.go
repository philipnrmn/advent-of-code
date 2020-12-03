package main

import . "../lib"

func part1(input []string) (int, error) {
	return countForSlope(input, 3, 1), nil
}

func part2(input []string) (int, error) {
	// (assume at least one count is nonzero)
	count := 1
	cases := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}

	for _, c := range cases {
		count *= countForSlope(input, c[0], c[1])
	}
	return count, nil
}

func countForSlope(input []string, sX, sY int) int {
	var count int

	width := len(input[0])
	x := 0
	for y := 0; y < len(input); y += sY {
		r := input[y][x%width]
		if r == '#' {
			count++
		}

		x += sX
	}
	return count
}

func main() {
	Solve(3, part1, part2)
}
