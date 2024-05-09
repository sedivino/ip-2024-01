package main

import "fmt"

func main() {
	// Leitura dos valores de entrada
	var valorIngresso, valorInicial, valorFinal float64
	fmt.Print("Digite o valor do ingresso: ")
	fmt.Scan(&valorIngresso)
	fmt.Print("Digite o valor inicial do intervalo: ")
	fmt.Scan(&valorInicial)
	fmt.Print("Digite o valor final do intervalo: ")
	fmt.Scan(&valorFinal)

	// Verifica se o intervalo é válido
	if valorInicial >= valorFinal {
		fmt.Println("INTERVALO INVALIDO.")
		return
	}

	// Variáveis para armazenar o melhor lucro e o número de ingressos correspondente
	var melhorLucro, melhorNumIngressos float64
	var melhorValorIngresso float64

	// Loop para testar cada valor de ingresso no intervalo especificado
	for valor := valorInicial; valor <= valorFinal; valor += 1.0 {
		// Calcula o número de ingressos vendidos para o valor atual
		numIngressos := calcularNumIngressos(valor, valorIngresso)
		// Calcula o lucro para o valor atual
		lucro := calcularLucro(numIngressos, valor)
		// Verifica se o lucro atual é maior que o melhor lucro encontrado até agora
		if lucro > melhorLucro {
			melhorLucro = lucro
			melhorNumIngressos = numIngressos
			melhorValorIngresso = valor
		}
		// Imprime o valor do ingresso, o número de ingressos vendidos e o lucro para o valor atual
		fmt.Printf("V: %.2f, N: %.0f, L: %.2f\n", valor, numIngressos, lucro)
	}

	// Imprime o resumo com o melhor valor final, o lucro máximo e o número de ingressos correspondente
	fmt.Printf("\nMelhor valor final: %.2f\n", melhorValorIngresso)
	fmt.Printf("Lucro: %.2f\n", melhorLucro)
	fmt.Printf("Numero de ingressos: %.0f\n", melhorNumIngressos)
}

// Função para calcular o número de ingressos vendidos para um determinado valor de ingresso
func calcularNumIngressos(valor, valorIngresso float64) float64 {
	if valor < valorIngresso {
		return 120 + (25 * (valorIngresso - valor))
	} else if valor > valorIngresso {
		return 120 - (30 * (valor - valorIngresso))
	} else {
		return 120
	}
}

// Função para calcular o lucro para um determinado número de ingressos
func calcularLucro(numIngressos, valor float64) float64 {
	despesasFixas := 200.0 + (0.05 * numIngressos)
	return (valor * numIngressos) - despesasFixas
}
