package main

import (
	"fmt"
)

func main() {
	// Declaração das variáveis
	var n int

	// Lendo a quantidade n de elementos do vetor:
	fmt.Print("Digite o número (n) de elementos do vetor: ")
	fmt.Scan(&n)

	// Criando e preenchendo um vetor com n elementos:
	v := make([]int, n)
	
	fmt.Printf("Digite os elementos do vetor: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Print(v)
}