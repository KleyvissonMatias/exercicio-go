package main

import (
	"testing"
)

func TestTotalVogaisASCII(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"Hello", 2},
		{"World", 1},
		{"Programação", 4},
		{"", 0},
	}

	for _, tc := range testCases {
		got := totalVogaisASCII(tc.input)
		if got != tc.want {
			t.Errorf("Palavra: %s, Esperado: %d Obtido: %d", tc.input, tc.want, got)
		}
	}
}
