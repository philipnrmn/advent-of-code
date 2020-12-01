package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi(os.Args[1])
	fmt.Println(solve1(i))
}

func solve1(num int) int {
	x := 0
	y := 0

	// layer is the size of the side of the current layer
	layer := 1
	// max is the highest value in the currnet layer
	max := 1

	for i := 1; i < num; i += 2 {
		if num < i*i {
			layer = i
			max = i * i
			break
		}
		x++
		y++
	}

	// use a cursor to track where we are on the grid
	cursor := max
	// walk left along the bottom
	for cursor = cursor; cursor > max-layer; cursor-- {
		if cursor == num {
			return distance(x, y)
		}
		x--
	}

	// walk up along the left-hand side
	for cursor = cursor; cursor > max-(layer*2)+1; cursor-- {
		if cursor == num {
			return distance(x, y)
		}
		y--
	}

	// walk right along the top
	for cursor = cursor; cursor > max-(layer*3)+2; cursor-- {
		if cursor == num {
			return distance(x, y)
		}
		x++
	}

	// walk down along the right-hand side
	for cursor = cursor; cursor > max-(layer*4)+3; cursor-- {
		if cursor == num {
			return distance(x, y)
		}
		y++
	}

	return distance(x, y)
}

// distance is the manhattan distance of x,y from 1,1
func distance(x, y int) int {
	ax := iabs(x)
	ay := iabs(y)

	return ax + ay
}

// iabs is math.Abs, but for ints not floats
func iabs(x int) int {
	if x >= 0 {
		return x
	}
	return 0 - x
}
