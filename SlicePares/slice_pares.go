package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Números pares: ", returnSlicePares(array))
}

func returnSlicePares(array []int) []int {
	var resultPar []int

	for _, v := range array {
		if v%2 == 0 {
			resultPar = append(resultPar, v)
		}
	}
	return resultPar
}
