package rs

import "testing"

func TestParseString(t *testing.T) {
	tests := []struct {
		_   struct{}
		s   string
		v0  string
		v1  string
		err bool
	}{
		{
			s:   "",
			err: true,
		},
		{
			s:   "1..10",
			v0:  "1",
			v1:  "10",
			err: false,
		},
		{
			s:   "-11..10",
			v0:  "-11",
			v1:  "10",
			err: false,
		},
		{
			s:   "10..-1",
			v0:  "10",
			v1:  "-1",
			err: false,
		},
	}

	for _, test := range tests {
		v0, v1, err := parseString(test.s)
		if !test.err && err != nil {
			t.Fatal(err)
		} else if test.err && err == nil {
			t.Fatalf("error is nil")
		} else if test.err {
			continue
		}

		if v0 != test.v0 {
			t.Fatalf("v0: %v, but %v", test.v0, v0)
		}
		if v1 != test.v1 {
			t.Fatalf("v1: %v, but %v", test.v1, v1)
		}
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		_   struct{}
		v   interface{}
		rb  int
		ra  int
		err bool
	}{
		{
			v:   "",
			err: true,
		},
		{
			v:  "1..10",
			rb: 1,
			ra: 10,
		},
		{
			v:  "1..1",
			rb: 1,
			ra: 1,
		},
		{
			v:  "10..-1",
			rb: 10,
			ra: -1,
		},
	}

	for _, test := range tests {
		rb, ra, err := parseInt(test.v)
		if !test.err && err != nil {
			t.Fatal(err)
		} else if test.err && err == nil {
			t.Fatalf("error is nil")
		} else if test.err {
			continue
		}

		if rb != test.rb {
			t.Fatalf("rb: %v, but %v", test.rb, rb)
		}
		if ra != test.ra {
			t.Fatalf("ra: %v, but %v", test.ra, ra)
		}
	}
}
