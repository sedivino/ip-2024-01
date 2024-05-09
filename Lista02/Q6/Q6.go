package main

import "fmt"

func main() {
    var n int
    fmt.Println("Digite o número de elementos da sequência:")
    fmt.Scanln(&n)

    sequence := make([]int, n)

    fmt.Println("Digite a sequência de números inteiros:")
    for i := 0; i < n; i++ {
        fmt.Scan(&sequence[i])
    }

    maxLength := findMaxLength(sequence)
    fmt.Printf("O comprimento do segmento crescente máximo é: %d\n", maxLength)
}

func findMaxLength(sequence []int) int {
    maxLength := 0
    currentLength := 1

    for i := 1; i < len(sequence); i++ {
        if sequence[i] > sequence[i-1] {
            currentLength++
        } else {
            if currentLength > maxLength {
                maxLength = currentLength
            }
            currentLength = 1
        }
    }

    if currentLength > maxLength {
        maxLength = currentLength
    }

    return maxLength - 1
}
