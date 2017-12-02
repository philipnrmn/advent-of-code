package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
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

func solve2(spreadsheet string) int {
	lines := strings.Split(spreadsheet, "\n")
	sum := 0
	for _, l := range lines {
		is := ltois(l)
		i1, i2 := findEDVs(is)
		sum += i1 / i2
	}
	return sum
}

// findEDVs finds the first pair of evenly divisible numbers in a slice
func findEDVs(nums []int) (int, int) {
	// sort so that we always divide by a smaller number
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	l := len(nums)
	for i := 0; i < l; i++ {
		bigger := nums[i]
		for j := i + 1; j < l; j++ {
			smaller := nums[j]
			if bigger%smaller == 0 {
				return bigger, smaller
			}
		}
	}
	// This shouldn't happen, force a panic
	return 0, 0
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
