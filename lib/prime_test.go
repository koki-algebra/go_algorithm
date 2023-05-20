package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n   int
		exp bool
	}{
		{n: 29, exp: true},
		{n: 71, exp: true},
		{n: 173, exp: true},
		{n: 541, exp: true},
		{n: 1021, exp: true},
		{n: 1223, exp: true},
		{n: 942, exp: false},
		{n: 1024, exp: false},
		{n: 1224, exp: false},
		{n: 2000, exp: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n=%d", test.n), func(t *testing.T) {
			assert.Equal(t, test.exp, IsPrime(test.n))
		})
	}
}

func TestGetPrimes(t *testing.T) {
	tests := []struct {
		n   int
		exp []int
	}{
		{
			n:   16,
			exp: []int{2, 3, 5, 7, 11, 13},
		},
		{
			n:   100,
			exp: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n=%d", test.n), func(t *testing.T) {
			got := GetPrimes(test.n)
			assert.Equal(t, test.exp, got)
		})
	}
}
