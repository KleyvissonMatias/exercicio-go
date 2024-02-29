package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Números pares: ", retornaSlicePares(array))
}

func retornaSlicePares(array []int) []int {
	resultPar := []int{}

	for _, v := range array {
		if v%2 == 0 {
			resultPar = append(resultPar, v)
		}
	}
	sort.Ints(resultPar)

	return resultPar
}
