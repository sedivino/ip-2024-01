package main

import (
	"fmt"
)

func main() {

	// Declaração de variáveis
	var n int

	for
	
	
	for {
		
		fmt.Print("Digite a quantidade de elementos do vetor: ")
		fmt.Scan(&n)
		
		if n != 0 {

			v := make([]int, n)

			for i := 0; i < n; i++ {
				fmt.Scan(&v[i])
			}
		} else {
			break
		}
	}
	
}