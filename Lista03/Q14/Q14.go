package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declaração de variáveis
	var (
		n			int // n é o número de elementos da sequência.
		mediana		float64
	)

	// Lendo e validando n
	fmt.Print("Digiente a quantidade de termos da sequência numérica: ")
	fmt.Scan(&n)

	for {
		if n < 0 || n > 1000000 {
			fmt.Print("Número inválido. Digite um número entre 0 e 1.000.000: \n")
			fmt.Scan(&n)
		} else {
			break
		}
	}

	// Criando, lendo e ordenando um vetor com n elementos:
	v := make([]int, n)
	
	fmt.Print("Digiente os elementos da sequência numérica: ")

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	sort.Ints(v)

	// Calculando a mediana:
	if n%2 == 0 {
		mediana = float64(v[n/2-1]+v[n/2]) / 2.0 //Média dos valores centrais
	} else {
		mediana = float64(v[n/2]) // Valor central
	}

	//Imprimindo a mediana: 
	fmt.Printf("Mediana: %.2f", mediana)
}