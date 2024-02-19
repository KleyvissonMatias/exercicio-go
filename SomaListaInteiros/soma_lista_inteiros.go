package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Resultado da soma: ", listTotalSum(array))
}

func listTotalSum(array []int) []int {
	var result int

	for _, v := range array {
		result += v
	}
	return []int{result}
}
