package main

import (
	"fmt"
)

func main() {
	
	var salario, salarioReajustado float64
	
	// Entrada de dados
	
	fmt.Println("Digite o salário do funcionário:")
	fmt.Scanln(&salario)

	// Verificação do salário e aplicação do reajuste
	
	if salario <= 300.00 {
		salarioReajustado = salario * 1.50 // Reajuste de 50% para salários até R$ 300,00
	} else {
		salarioReajustado = salario * 1.30 // Reajuste de 30% para salários acima de R$ 300,00
	}

	// Saída de dados
	
	fmt.Printf("O salário reajustado do funcionário é: R$ %.2f\n", salarioReajustado)
}
