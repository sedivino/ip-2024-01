package main

import "fmt"

func main() {
	for {
		// Lendo o tamanho do vetor
		var N int
		fmt.Scan(&N)

		// Verificando se o tamanho é zero
		if N == 0 {
			break
		}

		// Criando o vetor e lendo seus elementos
		vetor := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&vetor[i])
		}

		// Encontrando o maior elemento do vetor
		maiorElemento := encontrarMaiorElemento(vetor)

		// Calculando a frequência de cada valor até o maior elemento
		frequencia := calcularFrequencia(vetor, maiorElemento)

		// Imprimindo a frequência de cada valor
		for i := 0; i <= maiorElemento; i++ {
			fmt.Printf("(%d) %d\n", i, frequencia[i])
		}
	}
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

func calcularFrequencia(vetor []int, maior int) []int {
	frequencia := make([]int, maior+1)
	for _, num := range vetor {
		frequencia[num]++
	}
	return frequencia
}
