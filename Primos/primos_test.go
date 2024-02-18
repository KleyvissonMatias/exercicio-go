package main

import "testing"

func TestEhPrimo(t *testing.T) {
	testCases := []struct {
		num  int
		want bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
	}

	for _, tc := range testCases {
		got := ehPrimo(tc.num)
		if got != tc.want {
			t.Errorf("Número %d, Esperado: %t, Obtido: %t", tc.num, tc.want, got)
		}
	}
}
