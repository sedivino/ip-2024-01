package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declaração das variáveis
	var n int

	// Leitura do tamanho do vetor
	fmt.Print("Digite o tamanho da lista de números: ")
	fmt.Scan(&n)

	// Verificando o tamanho da lista:
	for {
		if n < 1 || n >= 1000 {
			fmt.Print("Número inválido. Digite um número entre 1 e 1000: ")
			fmt.Scan(&n)
		 } else {
			break
		}
	}

	// Criando e preenchendo a lista numérica
	v := make([]int, n)

	fmt.Print("Digite os elementos da lista: \n")

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	// Ordenando a lista:
	sort.Ints(v)

	// Imprimindo os elementos da lista
	for _, num := range v {
		fmt.Println(num)
	}

}