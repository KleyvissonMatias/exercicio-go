package main

import (
	"fmt"
	"sort"
)

func main() {
	elementos := []int{1, 1, 2, 2, 3, 4, 6, 5, 8, 9, 8}
	fmt.Println("Elementos: ", retornaElementosUnicos(elementos))
}

func retornaElementosUnicos(elementos []int) []int {
	unicosMap := make(map[int]bool)
	elementosUnicos := []int{}

	for _, elemento := range elementos {
		unicosMap[elemento] = true
	}

	for elemento := range unicosMap {
		elementosUnicos = append(elementosUnicos, elemento)
	}

	sort.Ints(elementosUnicos)

	return elementosUnicos
}
