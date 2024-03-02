package main

import "testing"

func TestSum(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{-1, 1}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3, 4, 5, -5, -4, -3, -2, -1}, 0},
	}

	for _, tc := range testCases {
		got := retornaSoma(tc.input)
		if got != tc.want {
			t.Errorf("Entrada: %v, Esperado: %d, Obtido: %d", tc.input, tc.want, got)
		}
	}
}
