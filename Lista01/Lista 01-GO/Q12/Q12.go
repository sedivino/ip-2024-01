package main

import (
	"fmt"
)

func main() {
	var HorasUtilizadas, ValorAPagar float64

	const TaxaBase = 10.0 // Taxa cobrada para as 3 primeiras horas
	const TaxaExtra = 5.0 // Taxa extra cobrada para cada hora adicional

	// Entrada de dados

	fmt.Println("Digite o número de horas que a charrete foi utilizada:")
	fmt.Scanln(&HorasUtilizadas)

	if HorasUtilizadas <= 3{
		ValorAPagar = TaxaBase
	} else {
		ValorAPagar = TaxaBase + (HorasUtilizadas - 3) * TaxaExtra
	}

	// Saída de dados

	fmt.Printf("O valor a pagar pelo cliente é R$ %.2f \n", ValorAPagar)
}