package main

import (
	"strings"

	. "../lib"
)

func part1(state []string) (int, error) {
	var stable bool
	for {
		state, stable = iterate(state)
		if stable {
			break
		}
	}

	occupied := 0
	for _, row := range state {
		occupied += strings.Count(row, "#")
	}

	return occupied, nil
}

func iterate(state []string) ([]string, bool) {
	stable := true
	newState := make([]string, len(state))
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			adj := countAdjacent(state, i, j)
			n := state[i][j]
			if n == 'L' && adj == 0 {
				n = '#'
				stable = false
			}
			if n == '#' && adj >= 4 {
				n = 'L'
				stable = false
			}
			newState[i] += string(n)
		}
	}
	return newState, stable
}

func countAdjacent(state []string, y, x int) int {
	adjacent := []byte{
		at(state, y-1, x-1), at(state, y-1, x), at(state, y-1, x+1),
		at(state, y, x-1), at(state, y, x+1),
		at(state, y+1, x-1), at(state, y+1, x), at(state, y+1, x+1),
	}
	count := 0
	for _, r := range adjacent {
		if r == '#' {
			count++
		}
	}
	return count
}

func at(state []string, y, x int) byte {
	if x < 0 || y < 0 || x >= len(state[0]) || y >= len(state) {
		return '.'
	}
	return state[y][x]
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func main() {
	Solve(11, part1, part2)
}
