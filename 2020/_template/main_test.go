package main

import "testing"

var i = []string{}

func Test_part1(t *testing.T) {
	_, err := part1(i)
	if err != nil {
		t.Error(err)
	}
}

func Test_part2(t *testing.T) {
	_, err := part2(i)
	if err != nil {
		t.Error(err)
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
