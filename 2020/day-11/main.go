package main

import (
	"strings"

	. "../lib"
)

type strategy func([]string, int, int) int

var tolerance = 4

func part1(state []string) (int, error) {
	var stable bool
	for {
		state, stable = iterate(state, countAdjacent, 4)
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

func iterate(state []string, count strategy, tolerance int) ([]string, bool) {
	stable := true
	newState := make([]string, len(state))
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			c := count(state, i, j)
			n := state[i][j]
			if n == 'L' && c == 0 {
				n = '#'
				stable = false
			}
			if n == '#' && c >= tolerance {
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

func countVisible(state []string, y, x int) int {
	visible := []byte{
		find(state, y, x, -1, -1), find(state, y, x, -1, 0), find(state, y, x, -1, 1),
		find(state, y, x, 0, -1), find(state, y, x, 0, 1),
		find(state, y, x, 1, -1), find(state, y, x, 1, 0), find(state, y, x, 1, 1),
	}
	count := 0
	for _, r := range visible {
		if r == '#' {
			count++
		}
	}
	return count

}

func find(state []string, y, x, vy, vx int) byte {
	for {
		x += vx
		y += vy
		if x < 0 || y < 0 || x >= len(state[0]) || y >= len(state) {
			return '.'
		}
		if state[y][x] != '.' {
			return state[y][x]
		}
	}
}

func at(state []string, y, x int) byte {
	if x < 0 || y < 0 || x >= len(state[0]) || y >= len(state) {
		return '.'
	}
	return state[y][x]
}

func part2(state []string) (int, error) {
	var stable bool
	for {
		state, stable = iterate(state, countVisible, 5)
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

func main() {
	Solve(11, part1, part2)
}
