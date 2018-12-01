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
	result := 0
	for _, field := range strings.Fields(input) {
		m := 1
		if field[0] == byte('-') {
			m = -1
		}
		n, _ := strconv.Atoi(field[1:])
		result += n * m

	}
	return result
}
