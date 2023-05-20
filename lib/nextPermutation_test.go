package lib

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	intSlice := []int{1, 2, 3}
	intExp := [][]int{
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	for i := 1; NextPermutation(sort.IntSlice(intSlice)); i++ {
		assert.Equal(t, intExp[i-1], intSlice)
	}

	stringSlice := []string{"aaa", "acb", "bbb"}
	stringExp := [][]string{
		{"aaa", "bbb", "acb"},
		{"acb", "aaa", "bbb"},
		{"acb", "bbb", "aaa"},
		{"bbb", "aaa", "acb"},
		{"bbb", "acb", "aaa"},
	}
	for i := 1; NextPermutation(sort.StringSlice(stringSlice)); i++ {
		assert.Equal(t, stringExp[i-1], stringSlice)
	}
}
