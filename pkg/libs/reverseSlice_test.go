package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseSlice(t *testing.T) {
	t.Parallel()

	func() {
		slice := []int{1, 2, 3, 4, 5}
		exp := []int{5, 4, 3, 2, 1}
		got := ReverseSlice(slice)
		assert.Equal(t, exp, got)
		assert.NotEqual(t, slice, exp)
	}()

	func() {
		slice := []string{"a", "b", "c", "d", "e"}
		exp := []string{"e", "d", "c", "b", "a"}
		got := ReverseSlice(slice)
		assert.Equal(t, exp, got)
		assert.NotEqual(t, slice, exp)
	}()

	func() {
		type test struct {
			name string
		}

		slice := []test{
			{name: "a"},
			{name: "b"},
			{name: "c"},
			{name: "d"},
			{name: "e"},
		}
		exp := []test{
			{name: "e"},
			{name: "d"},
			{name: "c"},
			{name: "b"},
			{name: "a"},
		}
		got := ReverseSlice(slice)
		assert.Equal(t, exp, got)
		assert.NotEqual(t, exp, slice)
	}()
}
