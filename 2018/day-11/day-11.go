package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) (int, int) {
	serial, _ := strconv.Atoi(input)

	bigX := 0
	bigY := 0
	top3x3 := 0
	grid := make([][]int, 300)
	for y := range grid {
		grid[y] = make([]int, 300)
		for x := range grid[y] {
			grid[y][x] = levelOf(x, y, serial)
		}
	}

	for y := range grid {
		for x := range grid[y] {
			sum := sumAt(grid, x, y)
			if sum > top3x3 {
				top3x3 = sum
				bigX = x
				bigY = y
			}
		}
	}

	fmt.Println("Top power level was", top3x3)
	return bigX, bigY
}

func levelOf(x, y, serial int) int {
	rackID := x + 10
	level := rackID * y
	level += serial
	level *= rackID
	level -= level % 100
	level /= 100
	level = level % 10
	return level - 5
}

func sumAt(grid [][]int, x, y int) int {
	result := 0

	for i := 0; i < 3; i++ {
		if y+i >= 300 {
			break
		}
		for j := 0; j < 3; j++ {
			if x+j >= 300 {
				break
			}
			result += grid[y+i][x+j]
		}
	}
	return result
}
