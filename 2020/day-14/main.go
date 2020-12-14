package main

import (
	"regexp"
	"strconv"

	. "../lib"
)

var re = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func part1(input []string) (int, error) {
	var result int
	rgst := map[int]int{}
	for _, c := range chunk(input) {
		mask, instr := parse(c)
		for _, i := range instr {
			adr, val := i[0], i[1]
			rgst[adr] = masked(val, mask)
		}
	}
	for _, v := range rgst {
		result += v
	}
	return result, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func chunk(input []string) [][]string {
	var result [][]string
	var chunk []string

	for _, line := range input {
		if line[0:4] == "mask" && len(chunk) > 0 {
			result = append(result, chunk)
			chunk = []string{}
		}
		chunk = append(chunk, line)
	}

	if len(chunk) > 0 {
		result = append(result, chunk)
	}

	return result
}

func parse(input []string) (string, [][]int) {
	mask := input[0][7:]
	var instructions [][]int
	for _, line := range input[1:] {
		groups := re.FindStringSubmatch(line)
		adr, _ := strconv.Atoi(groups[1])
		val, _ := strconv.Atoi(groups[2])
		instructions = append(instructions, []int{adr, val})
	}
	return mask, instructions
}

func masked(val int, mask string) int {
	var result int
	for i := 35; i >= 0; i-- {
		result = result << 1
		switch mask[35-i] {
		case 'X':
			result = result | val>>i&1
		case '1':
			result = result | 1
		}
	}
	return result
}

func main() {
	Solve(14, part1, part2)
}
