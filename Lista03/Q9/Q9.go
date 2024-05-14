package main

import (
	"fmt"
)

func main() {

	// Declaração de variáveis:
	const N = 3

	// Leitura dos pontos A, B e C:
	// Ponto A:
	a := [3]int{}
	
	fmt.Println("Digite as coordenadas do ponto A(x,y,z): ")
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	// Ponto B:
	b := [3]int{}
	
	fmt.Println("Digite as coordenadas do ponto B(x,y,z): ")
	for i := 0; i < N; i++ {
		fmt.Scan(&b[i])
	}
		
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print()
	
}