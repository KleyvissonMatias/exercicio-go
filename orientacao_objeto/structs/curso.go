package structs

import "fmt"

type Curso struct {
	Nome        string
	Periodos    int
	Disciplinas map[int][]string
}

func NewCurso() Curso {
	return Curso{}
}

func (c Curso) GetCursoPorNome(nomeCurso string) map[int][]string {
	cursos := c.getCursos()

	return cursos[nomeCurso]
}

func (c Curso) GetTotalPeriodosCurso(nomeCurso string) int {
	cursos := c.getCursos()

	return len(cursos[nomeCurso])
}

func (c Curso) getCursos() map[string]map[int][]string {
	return map[string]map[int][]string{
		"Ciência da Computação": {
			1: {"Introdução à Computação", "Algoritmos e Complexidade"},
			2: {"Estruturas de Dados", "Arquitetura de Computadores"},
		},
		"Direito": {
			1: {"Introdução ao Direito", "Direito Constitucional"},
			2: {"Direito Civil", "Direito Penal"},
		},
	}
}

func (c Curso) ToString() string {
	return fmt.Sprintf("Nome: %s, Disciplinas: %v, Periodos: %d", c.Nome, c.Disciplinas, c.Periodos)
}
