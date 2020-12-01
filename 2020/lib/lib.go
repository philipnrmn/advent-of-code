package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var NYI = errors.New("Not yet implemented")
var NOPE = errors.New("Did not find a solution")

type Solution func([]int) (int, error)

func Solve(day int, s ...Solution) {
	if len(os.Args) < 2 {
		barf("Usage: go run day-<day>/main.go <part>")
	}

	var part int
	var err error
	if part, err = strconv.Atoi(os.Args[1]); err != nil {
		barf("Bad value for part", err)
	}

	var input []int
	filename := fmt.Sprintf("day-%02d/input", day)

	var file *os.File
	if file, err = os.Open(filename); err != nil {
		barf(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var n int
		if n, err = strconv.Atoi(line); err != nil {
			barf("A line in ", filename, " was not an int", err)
		}
		input = append(input, n)
	}

	fmt.Printf("Solving day %d part %d...\n", day, part)

	out, err := s[part-1](input)
	if err != nil {
		barf(err)
	}

	fmt.Println(out)
}

func barf(err ...interface{}) {
	fmt.Fprintln(os.Stderr, err...)
	os.Exit(1)
}