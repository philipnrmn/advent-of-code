package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(solve1(os.Args[1], 600, 160, -1))
}

const (
	X = 0
	Y = 1
)

func solve1(input string, gridWidth, gridHeight, iterations int) string {
	pattern := regexp.MustCompile("position=< ?(-?\\d+),  ?(-?\\d+)> velocity=< ?(-?\\d+),  ?(-?\\d+)>")
	lines := strings.Split(input, "\n")
	pos := make([][]int, len(lines))
	vlx := make([][]int, len(lines))
	for i, line := range lines {
		matches := pattern.FindStringSubmatch(line)
		pos[i] = []int{parseInt(matches[1]), parseInt(matches[2])}
		vlx[i] = []int{parseInt(matches[3]), parseInt(matches[4])}
	}
	if iterations == -1 {
		ic := 0
		for {
			ic++
			pos = iterate(pos, vlx, 0)
			fmt.Println("Generation: ", ic, " Grid: ", gridWidth, "x", gridHeight)
			grid, viz := printGrid(pos, 0, 0, gridWidth, gridHeight)
			if viz {
				fmt.Println(grid)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	pos = iterate(pos, vlx, iterations)

	grid, _ := printGrid(pos, 0, 0, gridWidth, gridHeight)
	return grid
}

func iterate(pos [][]int, vlx [][]int, count int) [][]int {
	for i := range pos {
		pos[i][X] += vlx[i][X]
		pos[i][Y] += vlx[i][Y]
	}
	if count > 0 {
		return iterate(pos, vlx, count-1)
	}
	return pos
}

func printGrid(pos [][]int, x, y, width, height int) (string, bool) {
	var result string
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	visible := false
	for _, p := range pos {
		if p[X] >= x && p[X] < x+width && p[Y] >= y && p[Y] < y+height {
			grid[-y+p[Y]][-x+p[X]] = "#"
			visible = true
		}
	}

	for _, row := range grid {
		result += fmt.Sprintln(strings.Join(row, ""))
	}

	return result, visible
}

func parseInt(input string) int {
	if result, err := strconv.Atoi(input); err == nil {
		return result
	}
	fmt.Errorf("Could not convert %q to integer\n", input)
	return -1
}
