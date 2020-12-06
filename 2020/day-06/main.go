package main

import (
	. "../lib"
)

func part1(input []string) (int, error) {
	var sum int
	chunks := Chunk(input)

	for _, chunk := range chunks {
		sum += len(count(chunk))
	}

	return sum, nil
}

func part2(input []string) (int, error) {
	var sum int
	chunks := Chunk(input)

	for _, chunk := range chunks {
		for _, i := range count(chunk) {
			if i == len(chunk) {
				sum++
			}
		}
	}

	return sum, nil
}

func count(chunk []string) map[rune]int {
	m := map[rune]int{}
	for _, line := range chunk {
		for _, r := range line {
			if _, ok := m[r]; ok {
				m[r] += 1
			} else {
				m[r] = 1
			}
		}
	}
	return m
}

func main() {
	Solve(06, part1, part2)
}
