package main

import (
	"strconv"

	. "../lib"
)

var compass = map[byte]int{'N': 0, 'E': 90, 'S': 180, 'W': 270}

func part1(input []string) (int, error) {
	var x, y, v int
	v = compass['E']
	for _, line := range input {
		d, n := parse(line)
		switch d {
		case 'L':
			v = turn(v, -n)
		case 'R':
			v = turn(v, n)
		case 'F':
			x, y = move(x, y, v, n)
		case 'N', 'E', 'S', 'W':
			x, y = move(x, y, compass[d], n)
		}
	}
	return abs(x) + abs(y), nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func parse(input string) (byte, int) {
	n, _ := strconv.Atoi(input[1:])
	return input[0], n
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func move(x, y, v, n int) (int, int) {
	switch v {
	case 0:
		y -= n
	case 90:
		x += n
	case 180:
		y += n
	case 270:
		x -= n
	}
	return x, y
}

func turn(v, n int) int {
	v += n
	if v < 0 {
		return 360 + v
	}
	if v >= 360 {
		return v - 360
	}
	return v
}

func main() {
	Solve(12, part1, part2)
}
