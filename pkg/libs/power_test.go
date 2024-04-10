package libs

import "testing"

func TestPow(t *testing.T) {
	tests := []struct {
		base, exponent, mod, expected int
	}{
		{2, 3, 5, 3},     // 2^3 % 5 = 3
		{3, 4, 7, 4},     // 3^4 % 7 = 4
		{5, 0, 10, 1},    // Any number^0 % mod = 1
		{2, 10, 100, 24}, // 2^10 % 100 = 24
	}

	for _, test := range tests {
		result := Pow(test.base, test.exponent, test.mod)
		if result != test.expected {
			t.Errorf("Pow(%d, %d, %d) = %d, want %d", test.base, test.exponent, test.mod, result, test.expected)
		}
	}
}
