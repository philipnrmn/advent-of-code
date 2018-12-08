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

type node struct {
	children []node
	metadata []int
}

func solve1(input string) int {
	fields := strings.Fields(input)
	nums := make([]int, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}

	root, _ := parse(nums)
	return sum(root)
}

func parse(input []int) (node, []int) {
	ccount := input[0]
	mdcount := input[1]
	result := node{
		children: make([]node, ccount),
		metadata: make([]int, mdcount),
	}

	remainder := input[2:]
	for i := 0; i < ccount; i++ {
		result.children[i], remainder = parse(remainder)
	}
	result.metadata = remainder[:mdcount]

	return result, remainder[mdcount:]
}

func sum(n node) int {
	result := 0
	for _, n := range n.children {
		result += sum(n)
	}
	for _, i := range n.metadata {
		result += i
	}
	return result
}
