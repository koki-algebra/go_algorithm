package libs

import (
	"testing"
)

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

func TestCountBinaryDigits(t *testing.T) {
	tests := []struct {
		input, expected int
	}{
		{0, 1},     // 0 in binary has 1 digits
		{1, 1},     // 1 in binary has 1 digit
		{10, 4},    // 10 in binary is 1010, which has 4 digits
		{255, 8},   // 255 in binary is 11111111, which has 8 digits
		{1024, 11}, // 1024 in binary is 10000000000, which has 11 digits
	}

	for _, test := range tests {
		result := countBinaryDigits(test.input)
		if result != test.expected {
			t.Errorf("countBinaryDigits(%d) = %d, want %d", test.input, result, test.expected)
		}
	}
}
