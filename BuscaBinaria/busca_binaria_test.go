package main

import (
	"testing"
)

func TestValidaBin(t *testing.T) {
	testCases := []struct {
		list []int
		item int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 20}, 20, 5},
		{[]int{1, 2, 3, 4, 5, 7000}, 7000, 5},
		{[]int{1, 2, 3, 4, 5, 7000}, 380000, -1},
	}

	for _, tc := range testCases {
		got := ValidaBin(tc.list, tc.item)
		if got != tc.want {
			t.Errorf("Lista %v e item %d, Esperado %d, Obtido %d", tc.list, tc.item, tc.want, got)
		}
	}
}
