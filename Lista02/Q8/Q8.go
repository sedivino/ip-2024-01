package main

import (
	"fmt"
)

func main() {

	var n int

	// Lendo a quantidade de times
	fmt.Printf("Digite a quantidade de times do campeonato: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Campeonato inválido!")
		return
	}

	final := 1

	fmt.Println("Finais possíveis:")

	// Loop para iteração de todos os times
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			
			fmt.Printf("Final %d: time%d X time%d\n", final, i, j)
			
			final++
		}		
	}	
}