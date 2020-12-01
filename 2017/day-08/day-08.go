package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	register  string
	amount    int
	condition condition
}

type condition struct {
	register string
	operator string
	value    int
}

// biggest tracks the largest value recorded by exec()
var biggest int

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
}

func solve1(input string) int {
	lines := strings.Split(input, "\n")
	registers := map[string]int{}
	for _, l := range lines {
		instr := parse(l)
		registers = exec(instr, registers)
	}

	big := 0
	for _, v := range registers {
		if v > big {
			big = v
		}
	}

	return big
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")
	registers := map[string]int{}
	for _, l := range lines {
		instr := parse(l)
		registers = exec(instr, registers)
	}

	return biggest
}

// parse returns an instruction from a string
func parse(s string) instruction {
	re := regexp.MustCompile(`(\w+) (\w+) (-?\d+) if (\w+) ([<>=!]+) (-?\d+)`)
	sm := re.FindAllStringSubmatch(s, -1)[0][1:]
	amt, _ := strconv.Atoi(sm[2])
	val, _ := strconv.Atoi(sm[5])
	return instruction{
		register:  sm[0],
		operation: sm[1],
		amount:    amt,
		condition: condition{
			register: sm[3],
			operator: sm[4],
			value:    val,
		},
	}
}

// exec executes the instruction against the registers
func exec(i instruction, r map[string]int) map[string]int {
	if !eval(i.condition, r) {
		return r
	}

	switch i.operation {
	case "inc":
		r[i.register] += i.amount
	case "dec":
		r[i.register] -= i.amount
	}

	// specific to part 2
	if r[i.register] > biggest {
		biggest = r[i.register]
	}

	return r
}

// eval is true if the condition is true, given the registers
func eval(c condition, r map[string]int) bool {
	i := r[c.register]
	switch c.operator {
	case ">":
		return i > c.value
	case ">=":
		return i >= c.value
	case "<":
		return i < c.value
	case "<=":
		return i <= c.value
	case "==":
		return i == c.value
	case "!=":
		return i != c.value
	}
	fmt.Println("Error: unrecognized condition operator", c.operator)
	return false
}
