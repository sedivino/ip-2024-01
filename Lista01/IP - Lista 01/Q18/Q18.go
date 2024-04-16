package main

import "fmt"

func main() {
    var valorInicial, razao, numeroElementos int

    // Entrada de dados
    
	fmt.Println("Digite o valor inicial, a razão e o número de elementos separados por espaço:")
    fmt.Scanln(&valorInicial, &razao, &numeroElementos)

    // Cálculo da soma da progressão aritmética
    
	soma := ((2 * valorInicial) + (razao * (numeroElementos - 1)) * numeroElementos) / 2

    // Saída de dados
    
	fmt.Printf("A soma da progressão aritmética é: %d\n", soma)
}
