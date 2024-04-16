package main

import (
	"fmt"
)

func main() {
    var TF, TC, Ch float64

    // Entrada das temperaturas em Fahrenheit e da quantidade de chuva
    fmt.Println("Digite a temperatura em Fahrenheit:")
    fmt.Scanln(&TF)

    fmt.Println("Digite a quantidade de chuva:")
    fmt.Scanln(&Ch)

    // Conversão da Temperatura em Celsius
    TC = (5*TF - 160) / 9

    // Conversão da quantidade de chuva
    Ch = Ch * 25.4

    // Saída das conversões
    fmt.Println("O VALOR EM CELSIUS =", TC)
    fmt.Println("A QUANTIDADE DE CHUVA E =", Ch)
}
