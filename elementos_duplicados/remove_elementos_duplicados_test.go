package main

import (
	"reflect"
	"testing"
)

func TestRetornaElementosUnicos(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 2, 3, 4, 4, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4}, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		result := retornaElementosUnicos(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Para entrada %v, esperava-se %v, mas obteve %v", test.input, test.expected, result)
		}
	}
}
