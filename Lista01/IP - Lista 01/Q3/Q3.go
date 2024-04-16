package main

import (
	"fmt"
)

func main() {
    var n1, n2, n3, numeroConcatenado, quadrado int

    // Entrada de dados
    fmt.Print("Digite o primeiro número (centena): ")
    fmt.Scanln(&n1)

    fmt.Print("Digite o segundo número (dezena): ")
    fmt.Scanln(&n2)

    fmt.Print("Digite o terceiro número (unidade): ")
    fmt.Scanln(&n3)

    // Verificação de dígitos inválidos
    if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
        fmt.Println("DIGITO INVALIDO")
    } else {
        // Cálculo da concatenação
        numeroConcatenado = n1*100 + n2*10 + n3

        // Verificação de zeros à esquerda
        if n1 == 0 && n2 == 0 && n3 == 0 {
            fmt.Println("0")
        } else {
            fmt.Println(numeroConcatenado)
        }

        // Cálculo do quadrado
        quadrado = numeroConcatenado * numeroConcatenado
        fmt.Println("O quadrado do número é:", quadrado)
    }
}
