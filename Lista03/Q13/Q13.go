package main

import "fmt"

func main() {
	//Definição de variável
	var n int
	
	// Leitura e validação da quantidade n da sequência numérica:
	fmt.Print("Digite a quantidade de elementos da sequência: ")
	fmt.Scan(&n)

	for {
		if n <= 1 || n >= 1000000 {
			fmt.Print("Número inválido! Digite um número entre 1 e 1000000: ")
			fmt.Scan(&n)
		} else {
			break
		}
	}

	// Criando e preenchendo a lista numérica:
	v := make([]int, n)
	
	fmt.Print("Digite os elementos da lista numérica: ")

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	// Criar um mapa para contar a frequência de cada elemento
	frequencia := make(map[int]int)

	// Percorrer o vetor e contar a frequência de cada elemento
	for _, elemento := range v {
		frequencia[elemento]++
	}

	// Encontrar o elemento que mais se repete
	elementoMaisFrequente := v[0]
	maiorFrequencia := frequencia[v[0]]
	for _, elemento := range v {
		if frequencia[elemento] > maiorFrequencia {
			elementoMaisFrequente = elemento
			maiorFrequencia = frequencia[elemento]
		}
	}

	// Imprimir o elemento que mais se repete
	fmt.Printf("O elemento que mais se repete é %d, que aparece %d vezes.\n", elementoMaisFrequente, maiorFrequencia)
}
