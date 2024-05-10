package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N)

	// Declarando e lendo os elementos do vetor V
	V := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&V[i])
	}

	fmt.Scan(&M)

	// Realizando as buscas
	for i := 0; i < M; i++ {
		var num int
		fmt.Scan(&num)

		// Verificando se o número está no vetor V
		if busca(V, num) {
			fmt.Println("ACHEI")
		} else {
			fmt.Println("NAO ACHEI")
		}
	}
}

// Função para buscar um número em um vetor
func busca(V []int, num int) bool {
	for _, v := range V {
		if v == num {
			return true
		}
	}
	return false
}
