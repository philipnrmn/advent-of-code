package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) int {
	si := strings.Split(input, "\n")
	instrs := make([]int, len(si))
	for i, s := range si {
		instrs[i], _ = strconv.Atoi(s)
	}

	done := false
	cursor := 0

	var steps int
	for steps = 0; !done; steps++ {
		cursor, done = step(instrs, cursor)
	}
	return steps
}

func step(input []int, cursor int) (int, bool) {
	instr := input[cursor]
	input[cursor]++
	ninstr := instr + cursor
	return ninstr, ninstr >= len(input)
}
