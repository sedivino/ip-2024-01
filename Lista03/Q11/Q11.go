package main

import (
	"fmt"
)

func main() {
	//Declaração de variáveis
	var n, maxV, minW int

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
	
	// Criando e preenchendo um vetor V de n elementos:
	v := make([]int, n)
	maxV = 0
	
	fmt.Print("Digite os elementos do vetor: \n")
	
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])

		if v[i] > maxV {
			maxV = v[i]
		}
	}
	
	// Vamos definir o vetor inverso W e o preencher
	w := make([]int, n)
	
	for i := 0; i < len(v); i++ {
		minW = w[0]
		w[i]=v[len(v)-1-i]

		if w[i] < minW {
			minW = w[i]
		}
	}

	fmt.Println(v)
	fmt.Println(w)
	fmt.Println(maxV)
	fmt.Println(minW)
}