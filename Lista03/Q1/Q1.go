package main

import (
	"fmt"
)

func main() {

	// Entrada de n inteiros:
	var n, m int

	fmt.Print("Digite o tamanho do vetor: ")
	fmt.Scan(&n)

	// Criando um slice (vetor) com n inteiros
	v := make([]int, n)

	// For para preencher o vetor v
	fmt.Print("Digite os elementos do vetor: ")

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	// Entrada de M inteiros qtd de buscas
	fmt.Print("Digite a quantidade de buscas que deseja fazer: ")
	fmt.Scan(&m)

	// Criando um slice (vetor) com m inteiros
	v2 := make([]int, m)

	// For para preencher o vetor v2
	fmt.Print("Digite os elementos que deseja buscar: ")

	for i := 0; i < m; i++ {
		fmt.Scan(&v2[i])
	}
		
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if v2[i] == v[j] {
				fmt.Print("ACHEI! \n")
				break
			}
			
			if j == n - 1 {
				fmt.Print("NÃO ACHEI \n")
			}
		}

	}

}