package main

import "testing"

func TestPalavraMaisLonga(t *testing.T) {
	testCases := []struct {
		array []string
		want  string
	}{
		{[]string{"Amor", "Josi", "Kley", "KleyEJosi", "Paralelepipdo"}, "Paralelepipdo"},
		{[]string{"a", "bb", "ccc", "dddd"}, "dddd"},
		{[]string{"", "a", "ab", "abc"}, "abc"},
	}

	for _, tc := range testCases {
		got := palavraMaisLonga(tc.array)
		if got != tc.want {
			t.Errorf("Array: %v, Esperado: %s, Obtido: %s", tc.array, tc.want, got)
		}
	}
}
