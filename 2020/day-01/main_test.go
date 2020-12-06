package main

import "testing"

var i = []string{
	"1721",
	"979",
	"366",
	"299",
	"675",
	"1456",
}

func Test_part1(t *testing.T) {
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}

	if o != 514579 {
		t.Error("Wrong output")
	}
}

func Benchmark_part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part2(i)
	}
}
