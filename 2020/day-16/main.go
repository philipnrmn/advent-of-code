package main

import (
	"strings"

	. "../lib"
)

func part1(input []string) (int, error) {
	_, constraints, tickets := parse(input)

	valid := map[int]bool{}
	for _, c := range constraints {
		for i := c[0]; i <= c[1]; i++ {
			valid[i] = true
		}
		for i := c[2]; i <= c[3]; i++ {
			valid[i] = true
		}
	}

	rate := 0
	for _, t := range tickets[1:] {
		for _, f := range t {
			if !valid[f] {
				rate += f
			}
		}
	}
	return rate, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func parse(input []string) ([]string, [][]int, [][]int) {
	var labels []string
	var constraints [][]int

	chunks := Chunk(input)
	for _, line := range chunks[0] {
		p := strings.Split(line, ": ")
		labels = append(labels, p[0])
		q := strings.Split(p[1], " or ")
		r := strings.Split(q[0], "-")
		s := strings.Split(q[1], "-")
		constraints = append(constraints, []int{ToInt(r[0]), ToInt(r[1]), ToInt(s[0]), ToInt(s[1])})

	}

	mine := CsvInts(chunks[1][1:])
	nearby := CsvInts(chunks[2][1:])
	tickets := append(mine, nearby...)
	return labels, constraints, tickets
}

func main() {
	Solve(16, part1, part2)
}
