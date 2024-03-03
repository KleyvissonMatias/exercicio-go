package main

import (
	"fmt"
	"github.com/KleyvissonMatias/exercicio/orientacao_objeto/structs"
)

func main() {

	cursoCC := structs.NewCurso()
	cursoCC.Nome = "Ciência da Computação"
	cursoCC.Disciplinas = cursoCC.GetCursoPorNome("Direito")
	cursoCC.Periodos = cursoCC.GetTotalPeriodosCurso("Direito")

	aluno := structs.NewAluno()
	aluno.Nome = "Kleyvisson"
	aluno.Idade = 26
	aluno.Sexo = "Masculino"
	aluno.Matricula = "132234412"
	aluno.Contato = "81999136741"
	aluno.Email = "kleyvissonmatias@gmail.com"
	aluno.Curso = &cursoCC
	aluno.Periodo = 4

	coordenador := structs.NewCoordenador()
	coordenador.Nome = "Eduardo Calábria"
	coordenador.Idade = 50
	coordenador.Sexo = "Masculino"
	coordenador.Curso = &cursoCC
	coordenador.Contato = "8133345455"

	fmt.Println(aluno.ToString())
	fmt.Println(coordenador.ToString())
	fmt.Println(cursoCC.ToString())
}
