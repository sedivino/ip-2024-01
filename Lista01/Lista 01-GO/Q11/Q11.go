package main

import (
	"fmt"
)

func main() {
	var numero int

	// Entrada de dados

	fmt.Println("Digite um número inteiro:")
	fmt.Scanln(&numero)

	// Verificando a divisibilidade

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("O número é divisível por 3 e 5 ao mesmo tempo.")
	} else {
		fmt.Println("O número não é divisível por 3 e 5 ao mesmo tempo.")
	}
}
