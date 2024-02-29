package main

import (
	"reflect"
	"testing"
)

func TestRetornaElementosDuplicados(t *testing.T) {
	input := []int{1, 1, 2, 2, 3, 4, 6, 5, 8, 9, 8}
	expected := []int{1, 2, 8}

	result := retornaElementosDuplicados(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Para entrada %v, esperava-se %v, mas obteve %v", input, expected, result)
	}
}
