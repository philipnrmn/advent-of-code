package main

import "testing"

func Test_part1(t *testing.T) {
	i := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}

	if o != 514579 {
		t.Error("Wrong output")
	}
}
