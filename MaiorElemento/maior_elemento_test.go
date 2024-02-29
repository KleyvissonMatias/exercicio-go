package main

import (
	"reflect"
	"testing"
)

func TestMaiorSegundoMaiorElemento(t *testing.T) {
	tests := []struct {
		input  []int
		result [2]int
	}{
		{[]int{2, 7, 5, 12, 30, 20}, [2]int{30, 20}},
		{[]int{1, 2, 3, 4, 5}, [2]int{5, 4}},
		{[]int{9, 3, 6, 1, 8, 4, 5}, [2]int{9, 8}},
		{[]int{10, 10, 5, 5, 1}, [2]int{10, 5}},
	}

	for _, test := range tests {
		maior, segundoMaior := MaiorSegundoMaiorElemento(test.input)
		expected := test.result
		if !reflect.DeepEqual([2]int{maior, segundoMaior}, expected) {
			t.Errorf("Para a entrada %v, esperava-se %v, mas obteve %v", test.input, expected, [2]int{maior, segundoMaior})
		}
	}
}
