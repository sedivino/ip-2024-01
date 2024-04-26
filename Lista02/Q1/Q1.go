package main

import (
	"container/list"
	"fmt"
)

func main() {
	var {
		
		matricula int 
		prova, provas, lista, listas, trabalho presenca float64 = 0, 0, 0, 0, 0
		matricula
		notafinal[] float64
	}

	const qtdProvas = 8
	const qtdListas =5

	for i := 0; presenca != -1; i++ {

		fmt.Scan(&matricula)

		for i := 0; i < qtdProvas; i++ {
			fmt.Scan(&prova)
			provas += prova
		}

		for i := 0; i < qtdListas; i++ {
			fmt.Scan(&lista)
			listas += prova
		}

		fmt.Scan(&trabalho)

		fmt.Scan(&presenca)

		notafinal[i] = (.7 * provas/qtdProvas) + (.15 * listas/qtdListas) + (.15 * trabalho)

		fmt.Print(notafinal[i])
		presença = -1

	}
}