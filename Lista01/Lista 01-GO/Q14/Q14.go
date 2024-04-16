package main

import (
	"fmt"
)

func main() {
    var altura, aresta float64

    // Entrada de dados
    
	fmt.Println("Digite a altura da pirâmide em metros:")
    fmt.Scanln(&altura)

    fmt.Println("Digite a medida da aresta do hexágono da base em metros:")
    fmt.Scanln(&aresta)

    // Cálculo do volume da pirâmide
    
	areaBase := 3 * (3 * 1.732 * (aresta * aresta) / 2) // Fórmula para área de um hexágono regular
    volume := (areaBase * altura) / 3

    // Saída de dados
    
	fmt.Printf("O volume da pirâmide é = %.2f metros cúbicos\n", volume)
}
