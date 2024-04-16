package main

import "fmt"

func main() {
    var NC, n int

    // Entrada do número de conversões
    fmt.Println("Digite o número de conversões:")
    fmt.Scanln(&NC)

    n = 0
    for n < NC {
        var F, C float64

        // Entrada da temperatura em Fahrenheit
        fmt.Println("Digite a temperatura em Fahrenheit:")
        fmt.Scanln(&F)

        // Cálculo da temperatura em Celsius
        C = 5 * (F - 32) / 9

        // Saída da temperatura original e convertida
        fmt.Printf("Temperatura em Fahrenheit: %.2f Fahrenheit equivale a %.2f Celsius\n", F, C)

        n = n + 1
    }
}
