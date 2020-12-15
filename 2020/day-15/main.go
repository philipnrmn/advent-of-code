package main

import (
	"strconv"
	"strings"

	. "../lib"
)

func part1(input []string) (int, error) {
	turns := parse(input)
	index := map[int][]int{}
	for i, t := range turns {
		addIndex(&index, t, i)
	}

	for i := len(turns); i < 2020; i++ {
		occ := index[turns[len(turns)-1]]
		n := 0
		if len(occ) > 1 {
			n = abs(occ[len(occ)-1] - occ[len(occ)-2])
		}
		turns = append(turns, n)
		addIndex(&index, n, i)
	}
	return turns[len(turns)-1], nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func parse(input []string) []int {
	var result []int
	for _, i := range strings.Split(input[0], ",") {
		n, _ := strconv.Atoi(i)
		result = append(result, n)
	}

	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func addIndex(m *map[int][]int, v, i int) {
	if _, ok := (*m)[v]; ok {
		(*m)[v] = append((*m)[v], i)
	} else {
		(*m)[v] = []int{i}
	}
}

func main() {
	Solve(15, part1, part2)
}
