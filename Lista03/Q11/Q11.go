package main

import (
	"fmt"
)

func main() {
	//Declaração de variáveis
	var n int

	// Lendo o tamanho do vetor:
	fmt.Print("Digite o tamanho do vetor: ")
	fmt.Scan(&n)

	// Verificando o tamanho do vetor:
	for {
		if n < 1 || n >= 1000 {
			fmt.Print("Número inválido. Digite um número entre 1 e 1000: ")
			fmt.Scan(&n)
		} else {
			break
		}
	}
	
	// Criando e preenchendo um vetor de n elementos:
	v := make([]int, n)
	w := make([]int, n) // W é o vetor inverso.

	fmt.Print("Digite os elementos do vetor: \n")
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}
	
	fmt.Print(v)
	fmt.Print(w)
}