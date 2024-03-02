package main

import (
	"fmt"
	"math"
)

func main() {
	listaElementos := []int{2, 7, 5, 12, 30, 20}
	fmt.Println(retornaMaiorESegundoMaiorElemento(listaElementos))
}

func retornaMaiorESegundoMaiorElemento(listaElementos []int) (int, int) {
	maiorElemento := 0
	segundoMaiorElemento := math.MinInt

	for _, i := range listaElementos {
		if i > maiorElemento {
			segundoMaiorElemento = maiorElemento
			maiorElemento = i
		} else if i > segundoMaiorElemento && i != maiorElemento {
			segundoMaiorElemento = i
		}
	}

	return maiorElemento, segundoMaiorElemento
}
