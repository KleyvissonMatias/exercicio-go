package main

import (
	"testing"
)

func TestValidaNumeroSequencia(t *testing.T) {
	sequencia := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	tests := []struct {
		input    int
		expected bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, false},
		{8, true},
		{9, false},
		{10, false},
		{11, false},
		{12, false},
		{13, true},
		{14, false},
	}

	for _, test := range tests {
		result := validaNumeroSequencia(sequencia, test.input)
		if result != test.expected {
			t.Errorf("Entrada: %d, Esperado: %t, Obtido: %t", test.input, test.expected, result)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, test := range tests {
		result := fib(test.n)
		if result != test.expected {
			t.Errorf("Fib(%d) - Esperado: %d, Obtido: %d", test.n, test.expected, result)
		}
	}
}
