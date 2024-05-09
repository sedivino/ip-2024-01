package main

import (
	"fmt"
)

func main() {
	
	var (
		numero, somaPar, somaImpar, qtdPar, qtdImpar int
	)

	// Inicializando as variáveis de quantidade e soma:
	
	somaPar = 0;
	somaImpar = 0;
	qtdPar = 0;
	qtdImpar = 0;

	// Inicializando o loop para ler os dados:

	for {
		fmt.Println("Digite os valores da sequência. (Digite 0 para encerrar.)")
		fmt.Scanln(&numero)

		if numero == 0 {
			break 
		}

		if numero%2 == 0 {
			somaPar += numero	// Se o número for par a soma dos pares o recebe
			qtdPar++ 	// Incrementa a quantidade dos pares em +1
		} else {
			somaImpar += numero // Se o número não for par (ou seja, ímpar) a soma dos ímpares o recebe
			qtdImpar++	// Incrementa a quantidade dos ímpares em +1
		}
	}

	// Calculando as médias
	mediaPar := somaPar/(qtdPar)
	mediaImpar := somaImpar/(qtdImpar)
	
	// Imprimindo os resultados
	fmt.Printf("MEDIA PAR = %d\n", mediaPar)
	fmt.Printf("MEDIA ÍMPAR = %d\n", mediaImpar)

	} //fim