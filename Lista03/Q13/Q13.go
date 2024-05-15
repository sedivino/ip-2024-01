package main

import (
	"fmt"
)

func main() {

	// Declaração de variáveis
	var n int

	// Leitura e validação da quantidade n da sequência numérica:
	fmt.Print("Digite a quantidade de elementos da sequência: ")
	fmt.Scan(&n)

	for {
		if n <= 1 || n >= 1000000 {
			fmt.Print("Número inválido! Digite um número entre 1 e 1000000: ")
			fmt.Scan(&n)
		} else {
			break
		}
	}

	// Criando e preenchendo a lista numérica:
	v := make([]int, n)
	
	fmt.Print("Digite os elementos da lista numérica: ")

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	

	//fmt.Print(v)
}