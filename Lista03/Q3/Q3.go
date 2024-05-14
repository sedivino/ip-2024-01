package main

import (
	"fmt"
)

func main() {
	
	// Declaração das variáveis
	var n int

	// Criando um vetor v com n elementos:
	fmt.Print("Digite o tamanho do vetor: " )
	fmt.Scan(&n)

	v := make([]int, n)
	
	// Preenchendo o vetor:
	fmt.Print("Digite os elementos do vetor:\n")
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	// Invertendo a ordem dos valores do vetor e imprimindo o resultado:
	inverterSlice(v)
	fmt.Println("Vetor invertido: ", v)
}

func inverterSlice(slice []int) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		slice[i], slice[j] = slice[j], slice[i]
	}
}