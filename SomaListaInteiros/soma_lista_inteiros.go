package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Resultado da soma: ", listSum(array))
}

func listSum(array []int) []int {

	var result int
	var resultSum []int

	for _, v := range array {
		result += v
	}
	return append(resultSum, result)
}
