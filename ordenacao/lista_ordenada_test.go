package main

import (
	"reflect"
	"testing"
)

func TestListaOrdenada(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"Wsid", "Kley", "Aler", "Kade"}, []string{"Aler", "Kade", "Kley", "Wsid"}},
		{[]string{"banana", "apple", "orange"}, []string{"apple", "banana", "orange"}},
		{[]string{"3", "2", "1"}, []string{"1", "2", "3"}},
	}

	for _, test := range tests {
		result := listaOrdenada(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Entrada: %v, Esperado: %v, Obtido: %v", test.input, test.expected, result)
		}
	}
}
