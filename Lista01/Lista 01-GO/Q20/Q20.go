package main

import "fmt"

func main() {
    var horas, minutos, segundos int

    // Entrada de dados
    
	fmt.Println("Digite o tempo expresso em horas, minutos e segundos (HH:MM:SS):")
    fmt.Scanln(&horas, &minutos, &segundos)

    // Conversão para segundos
    
	totalSegundos := horas*3600 + minutos*60 + segundos

    // Saída de dados
   
	fmt.Printf("O tempo total em segundos é: %d\n", totalSegundos)
}
