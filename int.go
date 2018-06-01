package rs

import "fmt"

func Int(v interface{}) []int {
	before, after, err := parseInt(v)
	if err != nil {
		panic(err)
	}
	c := count(before, after)
	ret := make([]int, c)
	s := 1
	if !compare(before, after) {
		s = -1
	}

	fmt.Println(before, after)
	for i := 0; i < c; i++ {
		ret[i] = before + (i * s)
	}
	return ret
}

func compare(before, after int) bool {
	return before < after
}

func count(before, after int) int {
	if !compare(before, after) {
		return before - after + 1
	}
	return after - before + 1
}
