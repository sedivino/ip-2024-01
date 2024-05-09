package main

import (
	"fmt"
	"math"
)

func main() {

	var num int 
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num) 

	if num <= 1 { 
		fmt.Println("Numero invalido!") 
		return
	}

	isPrime := true 

	// Verifica se o número é divisível por algum número além de 1 e ele mesmo
	// A iteração começa de 2 e vai até a raiz quadrada do número
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 { 
			isPrime = false 
			break           
		}
	}

	// Verifica se a variável isPrime ainda é true após o loop
	if isPrime {
		fmt.Println("PRIMO")
	} else {
		fmt.Println("NAO PRIMO")
	}
}
