package main

import (
	"sort"

	. "../lib"
)

var preamble_length = 25

func part1(raw []string) (int, error) {
	input := Ints(raw)
	for i := preamble_length; i < len(input); i++ {
		preamble := input[i-preamble_length : i]
		valid := false
		for j := 0; j < preamble_length; j++ {
			for k := j + 1; k < preamble_length; k++ {
				if preamble[j]+preamble[k] == input[i] {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}
		if !valid {
			return input[i], nil
		}
	}
	return 0, NOPE
}

func part2(raw []string) (int, error) {
	invalid, _ := part1(raw)
	input := Ints(raw)
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			candidate := input[i:j]
			if sum(candidate) == invalid {
				sort.Ints(candidate)
				return candidate[0] + candidate[len(candidate)-1], nil
			}
		}
	}
	return 0, NOPE
}

func sum(s []int) int {
	if len(s) == 0 {
		return 0
	}
	result := s[0]
	for i := 1; i < len(s); i++ {
		result += s[i]
	}
	return result
}

func main() {
	Solve(9, part1, part2)
}
