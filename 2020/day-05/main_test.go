package main

import "testing"

var i = []string{
	"FBFBBFFRLR",
	"BFFFBBFRRR",
	"FFFBBBFRRR",
	"BBFFBBFRLL",
}

func Test_part1(t *testing.T) {
	e := 820
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}
	if o != e {
		t.Errorf("Expected %d, got %d", e, o)
	}
}

func Benchmark_part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part1(i)
	}
}

func Benchmark_part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part2(i)
	}
}
