package main

import "fmt"

func main() {
	array := []string{"Amor", "Josi", "Kley", "KleyEJosi", "Paralelepipdo"}
	fmt.Print("Palavra mais longa: ", palavraMaisLonga(array))
}

func palavraMaisLonga(array []string) string {
	palavra := ""
	for _, v := range array {
		if len(v) > len(palavra) {
			palavra = v
		}
	}
	return palavra
}
