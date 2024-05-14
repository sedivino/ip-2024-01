package main

import (
	"fmt"
)

func main() {
	
	// Declaração das variáveis
	var n int

	for {
		if n >= 0 {
			fmt.Print("Digite o número que deseja transformar para binário: ")
			fmt.Scan(&n)
			
			fmt.Printf("O número %d em binário é: %b \n", n, n)
		} else {
			break
		}
	}
}