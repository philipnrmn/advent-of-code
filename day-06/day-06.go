package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1], 10000))
}

type coord struct {
	cx int
	cy int
	id string
}

func solve1(input string) int {
	lines := strings.Split(input, "\n")
	coords := make([]coord, len(lines))
	for i, l := range lines {
		coords[i] = makeCoord(l)
	}

	top, right, bottom, left := boundsOf(coords)

	grid := make([][]string, bottom)
	for y := top; y < bottom; y++ {
		grid[y] = make([]string, right)
		for x := left; x < right; x++ {
			grid[y][x] = findNearest(coords, x, y)
		}
		// fmt.Println(grid[y])
	}

	most := 0
	for _, c := range coords {
		if c.cy == top || c.cx == left || c.cy == bottom || c.cx == right {
			continue
		}

		count := 0

		for y := top; y < bottom; y++ {
			for x := left; x < right; x++ {
				if grid[y][x] == c.id {
					count++
				}
			}
		}
		if count > most {
			most = count
		}
	}

	return most

}

func solve2(input string, limit int) int {
	lines := strings.Split(input, "\n")
	coords := make([]coord, len(lines))
	for i, l := range lines {
		coords[i] = makeCoord(l)
	}

	top, right, bottom, left := boundsOf(coords)

	grid := make([][]bool, bottom)
	for y := top; y < bottom; y++ {
		grid[y] = make([]bool, right)
		for x := left; x < right; x++ {
			grid[y][x] = distanceToAll(coords, x, y) < limit
		}
		// fmt.Println(grid[y])
	}

	area := 0
	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			if grid[y][x] {
				area++
			}
		}
	}

	return area

}

func makeCoord(line string) coord {
	xy := strings.Split(line, ", ")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return coord{cx: x, cy: y, id: fmt.Sprintf("%d,%d", x, y)}
}

func boundsOf(coords []coord) (int, int, int, int) {
	top := coords[0].cx
	right := coords[0].cy
	bottom := coords[0].cx
	left := coords[0].cy
	for _, i := range coords {
		switch {
		case i.cy < top:
			top = i.cy
		case i.cx > right:
			right = i.cx
		case i.cy > bottom:
			bottom = i.cy
		case i.cx < left:
			left = i.cx
		}
	}
	return top, right, bottom, left
}

func findNearest(coords []coord, x int, y int) string {
	var result string
	var nearest int
	var equidistant bool
	for i, c := range coords {
		distance := abs(c.cx-x) + abs(c.cy-y)
		switch {
		case i == 0:
			fallthrough
		case distance < nearest:

			equidistant = false
			nearest = distance
			result = c.id
		case distance == nearest:
			equidistant = true
		}
	}
	if equidistant {
		return "..."
	}
	return result
}

func distanceToAll(coords []coord, x int, y int) int {
	result := 0
	for _, c := range coords {
		distance := abs(c.cx-x) + abs(c.cy-y)
		result += distance
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
