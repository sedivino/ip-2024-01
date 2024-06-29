package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Ler a linha invertida da entrada
		invertedLine := scanner.Text()

		// Inverter a ordem das letras na linha
		reversed := reverseString(invertedLine)

		// Imprimir a linha com as letras na ordem correta
		fmt.Println(reversed)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada:", err)
	}
}

// Função para inverter a ordem das letras em uma string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
