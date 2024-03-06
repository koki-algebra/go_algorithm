package libs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		x, y int
		exp  int
	}{
		{x: 1, y: 0, exp: 1},
		{x: 15, y: 5, exp: 5},
		{x: 11, y: 55, exp: 11},
		{x: 1071, y: 1029, exp: 21},
		{x: 1000000000000000000, y: 999999999999999999, exp: 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(x,y)=(%d,%d)", test.x, test.y), func(t *testing.T) {
			t.Parallel()

			got := Gcd(test.x, test.y)
			assert.Equal(t, test.exp, got)
		})
	}
}
