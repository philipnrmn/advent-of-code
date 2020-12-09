package main

import (
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

func part2(input []string) (int, error) {
	return 0, NYI
}

func main() {
	Solve(9, part1, part2)
}
