package main

import "testing"

var i = []string{
	"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
	"byr:1937 iyr:2017 cid:147 hgt:183cm",
	"",
	"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
	"hcl:#cfa07d byr:1929",
	"",
	"hcl:#ae17e1 iyr:2013",
	"eyr:2024",
	"ecl:brn pid:760753108 byr:1931",
	"hgt:179cm",
	"",
	"hcl:#cfa07d eyr:2025 pid:166559648",
	"iyr:2011 ecl:brn hgt:59in",
}

func Test_part1(t *testing.T) {
	e := 2
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}
	if o != e {
		t.Errorf("Expected %d, got %d", e, o)
	}
}

func Test_part2(t *testing.T) {
	e := 0
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
