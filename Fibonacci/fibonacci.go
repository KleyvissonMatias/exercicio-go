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
		sequencia = append(sequencia, Fib(i))
	}

	fmt.Println("Digite um número Inteiro: ")
	var value int
	_, _ = fmt.Scanln(&value)
	fmt.Printf("%d é Fibonnaci? %s ", value, strconv.FormatBool(ValidaNumeroSequencia(sequencia, value)))
}

func Fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

func ValidaNumeroSequencia(seq []int, n int) bool {
	if slices.Contains(seq, n) {
		return true
	}
	return false
}
