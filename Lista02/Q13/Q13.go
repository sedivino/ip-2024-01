package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite o número de grãos para o primeiro quadro: ")
	fmt.Scan(&n)

	// Inicializa o total de grãos como 0
	totalGrãos := 0

	// Loop sobre cada quadrado do tabuleiro de xadrez (64 quadrados no total)
	for quadrado := 1; quadrado <= 64; quadrado++ {
		// Calcula o número de grãos para o quadrado atual
		grãosNoQuadrado := calcularGrãosNoQuadrado(n, quadrado)
		// Adiciona o número de grãos do quadrado atual ao total de grãos
		totalGrãos += grãosNoQuadrado
	}

	// Imprime o total de grãos no tabuleiro de xadrez
	fmt.Println("Total de grãos no tabuleiro de xadrez:", totalGrãos)
}

// Função para calcular o número de grãos para um quadrado específico do tabuleiro
func calcularGrãosNoQuadrado(n, quadrado int) int {
	// Se o quadrado for ímpar, é um quadrado escuro, e o número de grãos é o dobro de n
	if quadrado%2 == 1 {
		return n * 2
	}
	// Se o quadrado for par, é um quadrado claro, e o número de grãos é igual a n
	return n
}
