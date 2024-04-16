package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	// Entrada de dados
	fmt.Println("Digite o valor do coeficiente a:")
	fmt.Scanln(&a)

	fmt.Println("Digite o valor do coeficiente b:")
	fmt.Scanln(&b)

	fmt.Println("Digite o valor do coeficiente c:")
	fmt.Scanln(&c)

	// Cálculo do discriminante

	discriminante := b*b - 4*a*c

	// Saída de dados

	fmt.Printf("O valor do discriminante é %.2f \n", discriminante)
}
