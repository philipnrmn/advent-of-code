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

func solve1(spreadsheet string) int {
	lines := strings.Split(spreadsheet, "\n")
	sum := 0
	for _, l := range lines {
		is := ltois(l)
		sum += spread(is)
	}
	return sum
}

// spread finds the difference between the smallest and largest member of a
// slice of integers
func spread(is []int) int {
	lowest := is[0]
	highest := is[0]

	for _, i := range is {
		if i < lowest {
			lowest = i
		}
		if i > highest {
			highest = i
		}
	}
	return highest - lowest
}

// ltois converts a spreadsheet line to a slice of integers
func ltois(line string) []int {
	nums := strings.Split(line, "\t")
	var is []int
	for _, n := range nums {
		i, _ := strconv.Atoi(n)
		is = append(is, i)
	}
	return is
}
