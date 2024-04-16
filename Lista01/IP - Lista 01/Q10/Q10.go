package main

import (
	"fmt"
)

func main() {
	var a11, a12, a21, a22 float64

	// Entrada de dados

	fmt.Println("Digite o elemento a11 da matriz:")
	fmt.Scanln(&a11)

	fmt.Println("Digite o elemento a12 da matriz:")
	fmt.Scanln(&a12)

	fmt.Println("Digite o elemento a21 da matriz:")
	fmt.Scanln(&a21)

	fmt.Println("Digite o elemento a22 da matriz:")
	fmt.Scanln(&a22)

	// Calculando o determinante

	determinante := a11*a22 - a12*a21

	// Saída de dados

	fmt.Printf("O determinante da matriz é %.2f \n", determinante)

}