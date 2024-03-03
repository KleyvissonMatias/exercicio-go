package structs

import "fmt"

type Aluno struct {
	Nome      string
	Idade     int
	Sexo      string
	Matricula string
	Email     string
	Contato   string
	Curso     *Curso
	Periodo   int
}

func NewAluno() Aluno {
	return Aluno{}
}

func (a Aluno) ToString() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Sexo: %s, Matricula: %s, Email: %s,Contato: %s,"+
		"Curso: %s Periodo: %d",
		a.Nome, a.Idade, a.Sexo, a.Matricula, a.Email, a.Contato, a.Curso.Nome, a.Periodo)
}
