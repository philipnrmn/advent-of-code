package main

import (
	"fmt"
	"os"

	"container/ring"
	"strconv"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) string {
	g, _ := strconv.Atoi(input)
	elves := make([]*ring.Ring, 2)
	r := ring.New(2)

	elves[0] = r
	elves[0].Value = 3
	elves[1] = r.Next()
	elves[1].Value = 7

	printRing(r)

	for i := 0; i < g; i++ {
		fmt.Println(elves[0].Value)
		fmt.Println(elves[0].Value.(int))
		fmt.Println(elves[1].Value)
		fmt.Println(elves[1].Value.(int))
		sum := elves[0].Value.(int) + elves[1].Value.(int)
		if sum < 9 {
			s := ring.New(1)
			s.Value = sum
			r.Link(s)
			fmt.Println("inserting", s)
		} else {
			s := ring.New(2)
			units := sum % 10
			tens := (sum - units) / 10
			s.Value = tens
			s.Next()
			s.Value = units
			fmt.Println("inserting", s)
			r.Link(s)
		}
		for j := range elves {
			elves[j].Next()
			v := elves[j].Value.(int)
			for k := 0; k < v; k++ {
				elves[j].Next()
			}
		}
		printRing(r)
	}

	return "0"
}

func printRing(r *ring.Ring) {
	r.Do(func(p interface{}) {
		fmt.Printf("%d ", p.(int))
	})
	fmt.Println()
}
