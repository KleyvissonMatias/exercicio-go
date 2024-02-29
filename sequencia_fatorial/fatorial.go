package main

import "fmt"

func main() {
	fmt.Println("Digite um número Inteiro: ")
	var value int
	_, _ = fmt.Scanln(&value)
	fmt.Printf("Fatorial de %d! é: %d ", value, fat(value))
}

func fat(n int) int {
	if n == 0 {
		return 1
	}
	return n * fat(n-1)
}
