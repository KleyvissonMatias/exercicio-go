package structs

import "fmt"

type Professor struct {
	Nome  string
	Idade int
	Sexo  string
	Email string
	Curso *Curso
}

func (p Professor) ToString() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Sexo: %s, Curso: %s, Email: %s, Disciplinas: %v",
		p.Nome, p.Idade, p.Sexo, p.Curso.Nome, p.Email, p.Curso.Disciplinas)
}
