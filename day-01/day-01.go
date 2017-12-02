package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
}

func solve2(captcha string) int {
	l := len(captcha)
	offset := l / 2
	sum := 0

	for i := 0; i < l; i++ {
		c := getIntAt(captcha, i)
		cpp := getIntAt(captcha, (i+offset)%l)
		if c == cpp {
			sum += c
		}
	}
	return sum
}

func solve1(captcha string) int {
	l := len(captcha)

	first := getIntAt(captcha, 0)
	latest := first
	sum := 0

	for i := 1; i < l; i++ {
		c := getIntAt(captcha, i)
		if latest == c {
			sum += c
		}
		latest = c
	}

	if latest == first {
		sum += latest
	}

	return sum
}

// getIntAt is a helper which returns the integer value of the character
// at index i in string s
func getIntAt(s string, i int) int {
	n, _ := strconv.Atoi(string(s[i]))
	return n
}
