package main

import "fmt"

func main() {
    const custoAluminioMetroQuadrado = 100.00 // custo do alumínio por metro quadrado em reais

    var raio, altura float64

    // Entrada de dados
    fmt.Println("Digite o raio da lata em metros:")
    fmt.Scanln(&raio)

    fmt.Println("Digite a altura da lata em metros:")
    fmt.Scanln(&altura)

    // Cálculo da área lateral da lata (sem as bases)
    areaLateral := 2 * 3.14159 * raio * altura

    // Cálculo do custo total da lata
    custoTotal := custoAluminioMetroQuadrado * areaLateral

    // Saída de dados
    fmt.Printf("O VALOR DO CUSTO E = R$ %.2f\n", custoTotal)
}
