package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve1(os.Args[1:]))
}

func solve1(input []string) int {
	var banks []int
	for _, i := range input {
		b, _ := strconv.Atoi(i)
		banks = append(banks, b)
	}

	// this is easy to read, but probably slow
	results := map[string]bool{}
	for i := 0; ; i++ {
		h := hash(banks)
		if _, found := results[h]; found {
			return i
		}
		step(banks)
		results[h] = true
	}
}

func step(banks []int) {
	count := len(banks)
	// find the largest number in the banks
	big := 0
	ibig := 0
	for i, b := range banks {
		if b > big {
			ibig = i
			big = b
		}
	}
	// reset the biggest number to 0
	banks[ibig] = 0
	// redistribute it evenly
	for i := ibig + 1; big > 0; i++ {
		big--
		banks[i%count]++
	}
}

// hash returns a string (nonoptimal!) hash of memory state
func hash(banks []int) string {
	sbanks := make([]string, len(banks))
	for i, b := range banks {
		sbanks[i] = strconv.Itoa(b)
	}
	return strings.Join(sbanks, "-")
}
