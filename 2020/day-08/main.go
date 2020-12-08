package main

import (
	"fmt"
	"strconv"

	. "../lib"
)

func part1(input []string) (int, error) {
	var acc int
	visited := map[int]bool{}
	for i := 0; i < len(input); {
		if _, ok := visited[i]; ok {
			return acc, nil
		}
		visited[i] = true
		ins, val := parse(input[i])
		fmt.Printf("[%04d]\t%s\t%d\t%d\n", i, ins, val, acc)
		switch ins {
		case "nop":
			i++
			continue
		case "acc":
			acc += val
			i++
		case "jmp":
			i += val
		}
	}
	return 0, NOPE
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func parse(input string) (string, int) {
	ins := input[0:3]
	val, _ := strconv.Atoi(input[4:])
	return ins, val
}

func main() {
	Solve(8, part1, part2)
}
