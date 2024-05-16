package main

import (
	"fmt"
)

func main() {
	// Declaração de variáveis:
	var (
		t, n int // t: número de testes ; n: número de elementos do vetor.
	)

	// Lendo e validando a quantidade de testes:
	fmt.Print("Digite a quantidade de testes que serão feitos: \n")
	fmt.Scan(&t)

	for {
		if t < 1 || t > 10 {
			fmt.Print("Número de testes inválido; digite um número entre 1 e 10: ")
			fmt.Scan(&t)
		} else {
			break
		}
	}

	for i := 0; i < t; i++ {
		// Lendo e validando o número de elementos do vetor:
		fmt.Printf("Digite a quantidade de elementos do vetor %d: \n", i+1)
		fmt.Scan(&n)

		for {
			if n < 2 || n > 1000 {
				fmt.Print("Número de elementos inválido, digite um número entre 2 e 1000: \n")
				fmt.Scan(&n)	
			} else {
				break
			}
		}

		// Criando e preenchendo um vetor v com n elementos:
		v := make([]int, n)
		
		fmt.Printf("Digite os elementos da sequência numérica %d: \n", i+1)
		
		for j := 0; j < n; j++ {
			fmt.Scan(&v[j])
		}

				
		
	}

}