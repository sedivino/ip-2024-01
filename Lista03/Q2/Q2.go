package main

import (
	"fmt"
)

func main() {
	
	// Declaração das variáveis
	var n, K int

	// Criando um vetor v com n elementos:
	fmt.Print("Digite o tamanho do vetor: ")
	fmt.Scan(&n)

	// Verificando se o valor digitado é válido:
	for {
		if n < 1 || n > 1000 {
			fmt.Println("Número inválido. Digite um número entre 1 e 1000: ")
			fmt.Scan(&n)
		} else {
			break
		}
	}

	// Preenchendo o vetor:
	v := make([]int, n)
	fmt.Println("Digite os elementos do vetor: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	// Recebendo o número inteiro K para comparação:
	fmt.Println("Digite o número inteiro K para comparação: ")
	fmt.Scan(&K)

	// Comparando os elementos do vetor v com a constante K:
	count := 0 // Contador de elementos
	for _, valor := range v { 
		if valor >= K {
			count++
		}
	}

	// Imprimindo quantos valores do vetor V são maiores ou iguais a K:
	fmt.Printf("Há %d números que são maiores ou iguais a %d. ", count, K)

}