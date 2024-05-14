package main

import "fmt"

func main() {
	for {
		// Declaração de variáveis:
		var n int

		// Lendo o tamanho do vetor
		fmt.Print("Digite o tamanho do vetor: ")
		fmt.Scan(&n)

		// Verificando se o tamanho é zero
		if n == 0 {
			break
		}

		// Criando o vetor e lendo seus elementos
		v := make([]int, n)

		fmt.Print("Digite os elementos do vetor: ")

		for i := 0; i < n; i++ {
			fmt.Scan(&v[i])
		}

		// Encontrando o maior elemento e o índice da primeira ocorrência
		maiorElemento, indice := encontrarMaiorElemento(v)

		// Imprimindo o resultado
		fmt.Println(indice, maiorElemento)
	}
}

func encontrarMaiorElemento(vetor []int) (int, int) {
	// Inicializando o maior elemento e o índice da primeira ocorrência
	maiorElemento := vetor[0]
	indice := 0

	// Iterando pelo vetor para encontrar o maior elemento e seu índice
	for i, elemento := range vetor {
		if elemento > maiorElemento {
			maiorElemento = elemento
			indice = i
		}
	}

	return maiorElemento, indice
}
