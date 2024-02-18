package main

import "testing"

func TestFatorialSucesso(t *testing.T) {
	input := 2

	if Fat(input) == 0 {
		t.Error(Fat(input) == 0)
	}
}

func TestFatorialZeroSucesso(t *testing.T) {
	input := 0

	if Fat(input) < 0 {
		t.Error(Fat(input) < 0)
	}
}
