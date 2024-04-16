package main

import (
	"fmt"
)

func main() {
	var nota float64
	var conceito string

	// Entrada de dados

	fmt.Println("Digite a nota do aluno (entre 0 e 10):")
	fmt.Scanln(&nota)

	switch {
	case nota < 6:
		conceito = "D"
	
	case nota < 7.5:
		conceito = "C"

	case nota < 9:
		conceito = "B"

	default:
		conceito = "A"
	}

	// Saída de dados

	fmt.Printf("Nota = %.2f, Conceito = %s \n", nota, conceito)
}
