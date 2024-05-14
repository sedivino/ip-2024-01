package main

import (
	"fmt"
)

func main() {

	// Declaração de variáveis:
	var n int

	// Recebendo as notas
	fmt.Print("Digite a quantidade de notas: ")
	fmt.Scan(&n)

	// Verificando a quantidade de notas
	for {
		if n <= 1 || n >= 10000 {
			fmt.Print("Digite uma quantidade entre 1 e 10000: ")
			fmt.Scan(&n)
		} else {
			break
		}
	}

	// Recebendo as notas:
	notas := make([]int, n)

	fmt.Print("Digite as notas: ")
		
	for i := 0; i < n; i++ {
		fmt.Scan(&notas[i])
	}

	fmt.Print(notas)
}