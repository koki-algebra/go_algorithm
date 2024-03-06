package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		str string
		exp string
	}{
		"en": {
			str: "abcdefg",
			exp: "gfedcba",
		},
		"ja": {
			str: "あいうえお",
			exp: "おえういあ",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := ReverseString(test.str)
			assert.Equal(t, test.exp, got)
		})
	}
}
