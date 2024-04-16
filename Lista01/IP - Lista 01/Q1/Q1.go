package main

import (
    "fmt"
)

func main() {
    // Solicita ao usuário para inserir as três notas
    fmt.Println("Digite a primeira nota:")
    var nota1 float64
    fmt.Scanln(&nota1)

    fmt.Println("Digite a segunda nota:")
    var nota2 float64
    fmt.Scanln(&nota2)

    fmt.Println("Digite a terceira nota:")
    var nota3 float64
    fmt.Scanln(&nota3)

    // Calcula a média aritmética das três notas
    media := (nota1 + nota2 + nota3) / 3.0

    // Verifica se a média é maior ou igual a seis para decidir se o aluno foi aprovado ou reprovado
    if media >= 6.0 {
        fmt.Printf("Média: %.2f\nAPROVADO\n", media)
    } else {
        fmt.Printf("Média: %.2f\nREPROVADO\n", media)
    }
}

