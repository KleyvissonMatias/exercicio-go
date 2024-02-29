package main

import (
	"math/rand"
	"testing"
)

func TestMatrizDoisPorDois(t *testing.T) {
	rand.Intn(42)

	matriz := MatrizDoisPorDois()

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] < 0 || matriz[i][j] > 9 {
				t.Errorf("Elemento fora do intervalo [0, 9]: %d", matriz[i][j])
			}
		}
	}
}

func TestMatrizTresPorTres(t *testing.T) {
	rand.Intn(42)

	matriz := MatrizTresPorTres()

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] < 0 || matriz[i][j] > 9 {
				t.Errorf("Elemento fora do intervalo [0, 9]: %d", matriz[i][j])
			}
		}
	}
}
