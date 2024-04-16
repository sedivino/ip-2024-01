package main

import (
	"fmt"
)

func main() {
	var n int
	var soma float64

	// Entrada de dados
	
	fmt.Println("Digite um valor inteiro e positivo:")
	fmt.Scanln(&n)

	// Verificação se n é positivo
	
	if n <= 0 {
		fmt.Println("O valor deve ser inteiro e positivo.")
		return
	}

	// Cálculo da soma da série
	
	for i := 1; i <= n; i++ {
		soma += 1.0 / float64(i)
	}

	// Saída de dados
	fmt.Printf("A soma da série é: %.2f\n", soma)
}
