package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			got := Chmin(&a, test.b)
			assert.Equal(t, test.expBool, got)
			assert.Equal(t, test.expA, a)
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
			got := Chmax(&a, test.b)
			assert.Equal(t, test.expBool, got)
			assert.Equal(t, test.expA, a)
		})
	}
}
