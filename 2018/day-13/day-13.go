package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

const (
	LEFT     = -90
	STRAIGHT = 0
	RIGHT    = 90

	NORTH = 0
	EAST  = 90
	SOUTH = 180
	WEST  = 270

	OVERFLOW = 1000
)

type cart struct {
	x int
	y int
	v int
	t int
}

func (c *cart) turn(grid [][]rune) {
	c.v = (c.v + c.t) % 360
	c.t += 90
	if c.t > RIGHT {
		c.t = LEFT
	}
}

func (c *cart) tick(grid [][]rune) {
	switch c.v {
	case NORTH:
		c.y--
	case EAST:
		c.x++
	case SOUTH:
		c.y++
	case WEST:
		c.x--
	}

	switch grid[c.y][c.x] {
	case '/':
		switch c.v {
		case NORTH:
			c.v = EAST
		case EAST:
			c.v = NORTH
		case SOUTH:
			c.v = WEST
		case WEST:
			c.v = SOUTH
		}
	case '\\':
		switch c.v {
		case NORTH:
			c.v = WEST
		case EAST:
			c.v = SOUTH
		case SOUTH:
			c.v = EAST
		case WEST:
			c.v = NORTH
		}
	case '+':
		c.turn(grid)
	}
}

func solve1(input string) (int, int) {
	grid, carts := parseTrack(input)

	for i := 0; i < OVERFLOW; i++ {
		fmt.Println("\n\nGeneration", i)
		printTrack(grid, carts)
		time.Sleep(100 * time.Millisecond)
		for j := range carts {
			carts[j].tick(grid)
		}

		for y := range grid {
			for x := range grid[y] {
				cc := 0
				for _, c := range carts {
					if c.y == y && c.x == x {
						cc++
					}
					if cc > 1 {
						return x, y
					}
				}
			}
		}
	}

	fmt.Println("No collisions in", OVERFLOW, "iterations")
	return -1, -1
}

func parseTrack(track string) ([][]rune, []cart) {
	rows := strings.Split(track, "\n")
	grid := make([][]rune, len(rows))
	carts := []cart{}

	for y, row := range rows {
		grid[y] = make([]rune, len(row))
		for x, cell := range row {
			switch cell {
			case '^':
				grid[y][x] = '|'
				carts = append(carts, cart{x, y, NORTH, LEFT})
			case '>':
				grid[y][x] = '-'
				carts = append(carts, cart{x, y, EAST, LEFT})
			case 'v':
				grid[y][x] = '|'
				carts = append(carts, cart{x, y, SOUTH, LEFT})
			case '<':
				grid[y][x] = '-'
				carts = append(carts, cart{x, y, WEST, LEFT})
			default:
				grid[y][x] = cell
			}
		}
	}
	return grid, carts
}

func printTrack(grid [][]rune, carts []cart) {
	for y := range grid {
		for x := range grid[y] {
			cc := false
			for _, c := range carts {
				if c.x == x && c.y == y {
					cc = true
					switch c.v {
					case NORTH:
						fmt.Print("^")
					case EAST:
						fmt.Print(">")
					case SOUTH:
						fmt.Print("v")
					case WEST:
						fmt.Print("<")
					}
				}
			}
			if !cc {
				fmt.Print(string(grid[y][x]))
			}
		}
		fmt.Println()
	}
}
