package main

import "fmt"

func main() {
    var salarioMinimo float64
    var valorPorKW, valorTotalSemDesconto, valorTotalComDesconto float64
    var kWConsumidos int

    // Entrada de dados
    fmt.Print("Digite o valor do salário mínimo: ")
    fmt.Scanln(&salarioMinimo)

    fmt.Print("Digite a quantidade de kW consumidos pela residência: ")
    fmt.Scanln(&kWConsumidos)

    // Cálculo do valor por kW
    valorPorKW = (salarioMinimo * 0.7) / 100

    // Cálculo do valor total sem desconto
    valorTotalSemDesconto = valorPorKW * float64(kWConsumidos)

    // Cálculo do valor total com desconto de 10%
    valorTotalComDesconto = valorTotalSemDesconto * 0.9

    // Saída de dados
    fmt.Println("O valor em reais de cada kW é:", valorPorKW)
    fmt.Println("O valor total a ser pago pelo consumo da residência é:", valorTotalSemDesconto)
    fmt.Println("O novo valor a ser pago pela residência com desconto de 10% é:", valorTotalComDesconto)
}
