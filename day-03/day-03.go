package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	re := regexp.MustCompile("#(\\d)+ @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	for _, claim := range strings.Split(input, "\n") {
		matches := re.FindStringSubmatch(claim)
		x, _ := strconv.Atoi(matches[2])
		y, _ := strconv.Atoi(matches[3])
		w, _ := strconv.Atoi(matches[4])
		h, _ := strconv.Atoi(matches[5])

		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				grid[i][j]++
			}
		}
	}

	overlap := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] >= 2 {
				overlap++
			}
		}
	}
	return overlap

}
