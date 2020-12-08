package main

import (
	"fmt"
	"strconv"

	. "../lib"
)

func part1(input []string) (int, error) {
	acc, halted := halts(input, 0, 0)
	if halted {
		return acc, NOPE
	}
	return acc, nil
}

func part2(input []string) (int, error) {
	splicedInput := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		copy(splicedInput, input)
		ins, val := parse(input[i])
		switch ins {
		case "acc":
			continue
		case "nop":
			splicedInput[i] = fmt.Sprintf("jmp %d", val)
		case "jmp":
			splicedInput[i] = fmt.Sprintf("nop %d", val)
		}
		acc, halted := halts(splicedInput, 0, 0)
		if halted {
			return acc, nil
		}

	}
	return 0, NOPE
}

func halts(instructions []string, entrypoint int, acc int) (int, bool) {
	visited := map[int]bool{}
	for i := entrypoint; i < len(instructions); {
		if _, ok := visited[i]; ok {
			return acc, false
		}
		visited[i] = true
		ins, val := parse(instructions[i])
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
	return acc, true
}

func parse(input string) (string, int) {
	ins := input[0:3]
	val, _ := strconv.Atoi(input[4:])
	return ins, val
}

func main() {
	Solve(8, part1, part2)
}
