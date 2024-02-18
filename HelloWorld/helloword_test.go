package main

import (
	"strings"
	"testing"
)

func TestHelloSucesso(t *testing.T) {
	entrada := "Hello, World!"
	esperado := "Hello, World!"
	resultado := Hello()
	if !strings.EqualFold(resultado, entrada) {
		t.Errorf("Entrada: %s , Esperado: %s, Resultado: %s", entrada, esperado, resultado)
	}
}

func TestHelloFalha(t *testing.T) {
	entrada := "Hello, World!"
	esperado := "Hello!"
	resultado := Hello()
	if !strings.EqualFold(resultado, entrada) {
		t.Errorf("Entrada: %s , Esperado: %s, Resultado: %s", entrada, esperado, resultado)
	}
}
