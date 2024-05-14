package main

import (
	"fmt"
)

func main() {

	// Declaração de variáveis
	var n int

	// Lendo a quantidade n de elementos do vetor:
	fmt.Print("Digite o tamanho do vetor: ")
	fmt.Scan(&n)

	// Criando e preenchendo um vetor com n elementos:
	v := make([]int, n)

	// Definimos uma variável que receberá a soma dos elementos do vetor
	sum := 0

	fmt.Printf("Digite os elementos do vetor: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
		sum = sum + v[i]
	}

	fmt.Print(sum)
}