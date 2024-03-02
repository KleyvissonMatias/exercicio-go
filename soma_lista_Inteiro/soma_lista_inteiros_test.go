package main

import (
	"reflect"
	"testing"
)

func TestListTotalSum(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{15}},
		{[]int{0, 0, 0, 0, 0}, []int{0}},
		{[]int{1, -1}, []int{0}},
		{[]int{}, []int{0}},
	}

	for _, tc := range testCases {
		got := listTotalSum(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Entrada: %v, Esperado: %v, Obtido: %v", tc.input, tc.want, got)
		}
	}
}
