package main

import (
	"testing"
)

func TestEhPar(t *testing.T) {
	testCases := []struct {
		num  int
		want bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
	}

	for _, tc := range testCases {
		got := ehPar(tc.num)
		if got != tc.want {
			t.Errorf("Número: %d, Esperado: %t, Obtido: %t", tc.num, tc.want, got)
		}
	}
}
