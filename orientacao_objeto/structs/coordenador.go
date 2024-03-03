package structs

import "fmt"

type Coordenador struct {
	Nome    string
	Idade   int
	Sexo    string
	Contato string
	Curso   *Curso
}

func NewCoordenador() Coordenador {
	return Coordenador{}
}

func (c Coordenador) ToString() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Sexo: %s, Coordenador curso: %s, Contato: %s", c.Nome, c.Idade, c.Sexo,
		c.Curso.Nome, c.Contato)
}
