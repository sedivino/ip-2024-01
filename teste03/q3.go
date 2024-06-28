package main

import (
	"fmt"
)

func main() {
	var (
		n int
	)
	//Criando e preenchendo um vetor com n números inteiros:

	fmt.Print("Digite o tamanho do vetor: ")
	fmt.Scan(&n)

	v := make([]int, n)

	fmt.Print("Digite os elementos do vetor: ")

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println("Vetor original: ", v)

   //Invertendo o vetor v de números inteiros:
	bubbleSort(v)

	fmt.Println("Vetor ordenado: ", v)

}

func bubbleSort(vetor []int) {
	n := len(vetor)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if vetor[j] > vetor[j+1] {
				// Troca os elementos se estiverem fora de ordem
				vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
			}
		}
	}
}

