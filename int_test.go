package rs

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		before int
		after  int
		want   int
	}{
		{0, 0, 1},
		{0, 1, 2},
		{1, 0, 2},
		{-1, 1, 3},
		{1, -1, 3},
	}

	for _, test := range tests {
		r := count(test.before, test.after)

		if r != test.want {
			t.Fatalf("%v, %v", r, test.want)
		}
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		before int
		after  int
		want   bool
	}{
		{0, 0, false},
		{0, 1, true},
		{1, 0, false},
	}

	for _, test := range tests {
		r := compare(test.before, test.after)
		if r != test.want {
			t.Fatalf("%v, %v", r, test.want)
		}
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		v    interface{}
		want []int
	}{
		{
			"1..10",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"1..1",
			[]int{1},
		},
		{
			"-1..1",
			[]int{-1, 0, 1},
		},
		{
			"1..-1",
			[]int{1, 0, -1},
		},
	}

	for _, test := range tests {
		r := Int(test.v)
		if !reflect.DeepEqual(r, test.want) {
			t.Fatalf("want: %v, but %v", test.want, r)
		}
	}
}
