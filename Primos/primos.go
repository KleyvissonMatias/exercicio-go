package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Digite um número Inteiro: ")
	var value int
	_, _ = fmt.Scanln(&value)
	fmt.Printf("%d é primo? %s ", value, strconv.FormatBool(ehPrimo(value)))
}

func ehPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	if num%2 == 0 {
		return false
	}
	return true
}
