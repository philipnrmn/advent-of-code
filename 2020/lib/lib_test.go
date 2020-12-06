package lib

import (
	"reflect"
	"testing"
)

func Test_Chunk(t *testing.T) {
	i := []string{
		"1", "",
		"2", "3", "4", "",
		"",
		"5",
	}
	e := [][]string{
		[]string{"1"},
		[]string{"2", "3", "4"},
		[]string{"5"},
	}
	o := Chunk(i)
	if !reflect.DeepEqual(e, o) {
		t.Errorf("Expected %v, got %v", e, o)
	}
}
