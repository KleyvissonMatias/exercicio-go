package main

import (
	"reflect"
	"testing"
)

func TestReturnSlicePares(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8, 10}},
		{[]int{1, 3, 5, 7, 9}, []int{}},
		{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		got := retornaSlicePares(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Entrada: %v, Esperado: %v, Obtido: %v", tc.input, tc.want, got)
		}
	}
}
