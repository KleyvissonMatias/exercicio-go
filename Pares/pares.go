package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Digite um número Inteiro: ")
	var value int
	_, _ = fmt.Scanln(&value)
	fmt.Printf("%d é par? %s ", value, strconv.FormatBool(ehPar(value)))

}

func ehPar(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}
