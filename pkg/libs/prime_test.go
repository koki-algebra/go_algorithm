package libs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			assert.Equal(t, test.exp, IsPrime(test.n))
		})
	}
}

func TestGetPrimes(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			got := GetPrimes(test.n)
			assert.Equal(t, test.exp, got)
		})
	}
}

func TestPrimeFactorize(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		n   int
		exp map[int]int
	}{
		"123456789": {
			n: 123456789,
			exp: map[int]int{
				3:    2,
				3607: 1,
				3803: 1,
			},
		},
		"876543210": {
			n: 876543210,
			exp: map[int]int{
				2:    1,
				3:    2,
				5:    1,
				1997: 1,
				4877: 1,
			},
		},
		"987654321": {
			n: 987654321,
			exp: map[int]int{
				3:      2,
				17:     2,
				379721: 1,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := PrimeFactorize(test.n)
			assert.Equal(t, test.exp, got)
		})
	}
}
