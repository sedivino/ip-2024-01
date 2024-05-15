package main

import (
	"fmt"
)

func main() {

	// Declaração de variáveis:
	var n int

	// Recebendo as notas
	fmt.Print("Digite a quantidade de notas: ")
	fmt.Scan(&n)

	// Recebendo as notas e identificando o maior elemento
	notas := make([]int, n)
	
	fmt.Print("Digite as notas: ")
		
	for i := 0; i < n; i++ {
		fmt.Scan(&notas[i])
	}
	
	// Identificando o último elemento e fazendo sua contagem:
	ultimo := notas[len(notas)-1]
	contagem := 0 

	for _, elemento := range notas {
		if elemento == ultimo {
			contagem++
		}
	}

	notaMax := encontrarMaiorElemento(notas)


	// Imprimindo os resultados:
	fmt.Printf("Nota %d, %d vezes\n", ultimo, contagem)
	fmt.Printf("Nota %d, índice ", notaMax)
}

func encontrarMaiorElemento(vetor []int) int {
	maior := vetor[0]
	for _, num := range vetor {
		if num > maior {
			maior = num
		}
	}
	return maior
}