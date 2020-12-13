package main

import (
	"strconv"
	"strings"

	. "../lib"
)

func part1(input []string) (int, error) {
	ts, ids := parse(input)
	i := ts
	for {
		for _, id := range ids {
			if id == 0 {
				continue
			}
			if i%id == 0 {
				return (i - ts) * id, nil
			}
		}
		i++
	}
	return 0, NOPE
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func parse(input []string) (int, []int) {
	ts, _ := strconv.Atoi(input[0])
	var ids []int
	for _, r := range strings.Split(input[1], ",") {
		if r == "x" {
			ids = append(ids, 0)
		}
		id, _ := strconv.Atoi(r)
		ids = append(ids, id)
	}
	return ts, ids
}

func main() {
	Solve(13, part1, part2)
}
