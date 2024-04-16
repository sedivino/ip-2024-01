package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	// Entrada de dados
	fmt.Println("Digite um número inteiro entre 5 e 2000:")
	fmt.Scanln(&n)

	// Verificando se n está no intervalo válido
	if n < 5 || n > 2000 {
		fmt.Println("O número digitado está fora do intervalo válido.")
		return
	}

	// Gerando o quadrado de cada número par de 1 até n
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			quadrado := int(math.Pow(float64(i), 2))
			fmt.Printf("%d^2 = %d\n", i, quadrado)
		}
	}
}
