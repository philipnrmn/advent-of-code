package main

import "testing"

var i = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc",
}

func Test_part1(t *testing.T) {
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}

	if o != 2 {
		t.Errorf("Bad count: %d", o)
	}
}

func Test_part2(t *testing.T) {
	o, err := part2(i)
	if err != nil {
		t.Error(err)
	}
	if o != 1 {
		t.Errorf("Bad count: %d", o)
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
