package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	var sequencia []int

	n := 11
	for i := 1; i <= n; i++ {
		sequencia = append(sequencia, fib(i))
	}

	fmt.Println("Digite um número Inteiro: ")
	var value int
	_, _ = fmt.Scanln(&value)
	fmt.Printf("%d é Fibonnaci? %s ", value, strconv.FormatBool(validaNumeroSequencia(sequencia, value)))
}

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func validaNumeroSequencia(seq []int, n int) bool {
	if slices.Contains(seq, n) {
		return true
	}
	return false
}
