package main

import (
	"fmt"
	"sort"
)

func main() {
	lista := []string{"Wsid", "Kley", "Aler", "Kade"}
	fmt.Printf("Lista Ordenada: %v", listaOrdenada(lista))
}

func listaOrdenada(lista []string) []string {
	sort.Strings(lista)

	return lista
}
