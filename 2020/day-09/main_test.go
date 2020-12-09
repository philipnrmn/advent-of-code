package main

import "testing"

var i = []string{
	"35",
	"20",
	"15",
	"25",
	"47",
	"40",
	"62",
	"55",
	"65",
	"95",
	"102",
	"117",
	"150",
	"182",
	"127",
	"219",
	"299",
	"277",
	"309",
	"576",
}

func Test_part1(t *testing.T) {
	preamble_length = 5
	e := 127
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}
	if o != e {
		t.Errorf("Expected %d, got %d", e, o)
	}
}

func Test_part2(t *testing.T) {
	e := 62
	o, err := part2(i)
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
