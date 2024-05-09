package main

import (
	"fmt"
)

func main() {

	// Declaração das variáveis utlizadas
	var (
		matricula int
		horasTrabalhadas, valorHora float64
	)

	// Loop para leitura dos dados:
	for {
		fmt.Scan(&matricula, &horasTrabalhadas, &valorHora)

		if matricula == 0 {
			break
		}
		
		// Calculando o salário
		salario := horasTrabalhadas * valorHora

		fmt.Printf("%d %.2f\n", matricula, salario)
	}

}