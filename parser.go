package rs

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func parseInt(v interface{}) (rb, ra int, err error) {
	before, after, err := parse(v)
	if err != nil {
		return 0, 0, err
	}

	if rb, err = strconv.Atoi(before); err != nil {
		return 0, 0, fmt.Errorf("unsupported format: %v, %v, %v", v, before, reflect.TypeOf(before))
	}
	if ra, err = strconv.Atoi(after); err != nil {
		return 0, 0, fmt.Errorf("unsupported format: %v %v, %v", v, after, reflect.TypeOf(after))
	}
	return
}

func parse(v interface{}) (string, string, error) {
	switch v.(type) {
	case string:
		return parseString(v.(string))
	}

	return "", "", fmt.Errorf("unsupported format: %v", v)
}

func parseString(s string) (string, string, error) {
	if s == "" {
		return "", "", fmt.Errorf("input is empty string")
	}
	v := strings.Split(s, "..")
	if len(v) != 2 {
		return "", "", fmt.Errorf("unsupported format: %v", s)
	}
	return v[0], v[1], nil
}
