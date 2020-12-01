package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(solve1(os.Args[1], os.Args[2]))
}

func solve1(i1, i2 string) int {
	n, _ := strconv.Atoi(i1)
	hi, _ := strconv.Atoi(i2)

	elves := make([]int, n)
	elf := 0

	r := ring.New(1)
	r.Value = 0
	rr := r.Next()

	for i := 1; i <= hi; i++ {
		if i%23 == 0 {
			elves[elf] += i
			rr = rr.Prev().Prev().Prev().Prev().Prev().Prev().Prev().Prev()
			elves[elf] += rr.Value.(int)
			// fmt.Printf("Elf %d gets %d+%d\n", elf+1, i, rr.Value.(int))
			rr = rr.Prev()
			rr.Unlink(1)
			rr = rr.Next().Next()
		} else {
			s := ring.New(1)
			s.Value = i
			rr.Link(s)
			rr = rr.Next().Next()
		}

		// fmt.Printf("[%d] ", elf+1)
		// printRing(r)
		elf = (elf + 1) % n
	}

	highest := 0
	hiElf := 0
	for i, e := range elves {
		if e > highest {
			hiElf = i
			highest = e
		}
	}

	// fmt.Println(elves)

	fmt.Printf("Elf %d won with %d\n", hiElf+1, highest)
	return highest
}

func printRing(r *ring.Ring) {
	r.Do(func(p interface{}) {
		fmt.Printf("%d ", p.(int))
	})
	fmt.Println()
}
