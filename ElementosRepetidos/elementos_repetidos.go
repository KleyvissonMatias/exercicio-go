package main

import (
	"fmt"
	"sort"
)

func main() {
	elementos := []int{1, 1, 2, 2, 3, 4, 6, 5, 8, 9, 8}
	fmt.Println("Elementos duplicados: ", retornaElementosDuplicados(elementos))
}

func retornaElementosDuplicados(elementos []int) []int {
	elementosDuplicados := []int{}
	ocorrencias := make(map[int]int)

	for _, elemento := range elementos {
		ocorrencias[elemento]++
	}

	for elemento, ocorrencia := range ocorrencias {
		if ocorrencia > 1 {
			elementosDuplicados = append(elementosDuplicados, elemento)
		}
	}

	sort.Ints(elementosDuplicados)

	return elementosDuplicados
}
