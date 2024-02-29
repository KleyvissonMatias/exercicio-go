package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Matriz 2x2: ", MatrizDoisPorDois())
	fmt.Println()
	fmt.Print("Matriz 3x3: ", MatrizTresPorTres())
}

func MatrizDoisPorDois() [2][2]int {
	var matriz [2][2]int
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			matriz[i][j] = rand.Intn(10)
		}
	}
	return matriz
}

func MatrizTresPorTres() [3][3]int {
	var matriz [3][3]int
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			matriz[i][j] = rand.Intn(10)
		}
	}
	return matriz
}
