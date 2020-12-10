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

func part2(raw []string) (int, error) {
	input := Ints(raw)
	sort.Ints(input)
	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)
	branches := make([]int, len(input))
	branches[0] = 1
	branches[1] = 1

	for i := 2; i < len(input); i++ {
		branches[i] = branches[i-1]
		if input[i]-input[i-2] <= 3 {
			branches[i] += branches[i-2]
		}
		if i >= 3 && input[i]-input[i-3] <= 3 {
			branches[i] += branches[i-3]
		}
	}

	return branches[len(branches)-1], nil
}

func main() {
	Solve(10, part1, part2)
}
