package main

import (
	"fmt"
	"math"
)

func main() {
	var n, erro float64

	// Leitura do número e do erro
	fmt.Print("Digite o número cuja raiz quadrada deseja obter: ")
	fmt.Scan(&n)
	fmt.Print("Digite o erro desejado: ")
	fmt.Scan(&erro)

	// Chama a função para calcular a raiz quadrada usando o método de Newton-Raphson
	calculaRaizQuadrada(n, erro)
}

func calculaRaizQuadrada(n, erro float64) {
	// Chute inicial para a raiz quadrada
	x := n / 2.0
	// Variável para armazenar o valor anterior de x
	prevX := x

	// Iterações para encontrar a raiz quadrada com o erro desejado
	for {
		// Fórmula do método de Newton-Raphson para atualizar o valor de x
		x = 0.5 * (x + n/x)
		// Calcula o erro absoluto entre o valor atual de x e o valor anterior de x
		erroAtual := math.Abs(x - prevX)

		// Imprime o valor aproximado da raiz quadrada e o erro
		fmt.Printf("%.9f %.9f\n", x, erroAtual)

		// Verifica se o erro é menor ou igual ao erro desejado
		if erroAtual <= erro {
			break // Se o erro for menor ou igual ao erro desejado, encerra o loop
		}

		// Atualiza o valor anterior de x para o próximo ciclo de iteração
		prevX = x
	}
}
