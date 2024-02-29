package main

import "testing"

func TestFat(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, test := range tests {
		result := fat(test.input)
		if result != test.expected {
			t.Errorf("Fat(%d) Retornado %d, Esperado %d", test.input, result, test.expected)
		}
	}
}
