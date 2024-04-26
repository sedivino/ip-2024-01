package main

import (
	"fmt"
)

func calcularFatorial(n int) int {
	if n < 0 {
		fmt.Println("O número digitado é negativo")
		//return -1 // Retornar -1 para indicar um erro se o número for negativo
	} else if n == 0 {
		return 1 // Fatorial de 0 é 1
	}

	fatorial := 1
	for i := 1; i <= n; i++ {
		fatorial *= i
	}
	return fatorial
}

func main() {
	var num int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)

	resultado := calcularFatorial(num)
	fmt.Printf("O fatorial de %d é %d\n", num, resultado)
}
