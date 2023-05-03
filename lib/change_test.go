package lib

import "testing"

func TestChmin(t *testing.T) {
	tests := map[string]struct {
		a, b    int
		expA    int
		expBool bool
	}{
		"a is smaller than b": {a: 1, b: 2, expA: 1, expBool: false},
		"a is larger than b":  {a: 2, b: 1, expA: 1, expBool: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.a
			if got := Chmin(&a, test.b); got != test.expBool {
				t.Errorf("expected value is %t, but got %t", test.expBool, got)
			}
			if a != test.expA {
				t.Errorf("expected value of a is %d, but got %d", test.expA, a)
			}
		})
	}
}

func TestChmax(t *testing.T) {
	tests := map[string]struct {
		a, b    int
		expA    int
		expBool bool
	}{
		"a is smaller than b": {a: 1, b: 2, expA: 2, expBool: true},
		"a is larger than b":  {a: 2, b: 1, expA: 2, expBool: false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.a
			if got := Chmax(&a, test.b); got != test.expBool {
				t.Errorf("expected value is %t, but got %t", test.expBool, got)
			}
			if a != test.expA {
				t.Errorf("expected value of a is %d, but got %d", test.expA, a)
			}
		})
	}
}
