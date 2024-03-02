package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Digite uma palavra: ")
	var value string
	_, _ = fmt.Scanln(&value)
	fmt.Printf("Vogais: %d ", totalVogaisASCII(value))
}

func totalVogaisASCII(palavra string) int {
	total := 0
	listVogais := []int{65, 69, 73, 79, 85, 97, 101, 105, 111, 117}

	for _, char := range palavra {
		if slices.Contains(listVogais, int(char)) {
			total++
		}
	}
	return total
}
