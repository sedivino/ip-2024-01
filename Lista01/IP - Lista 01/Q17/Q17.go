package main

import "fmt"

func main() {
    
	var x, y int

    // Entrada de dados
   
	fmt.Println("Digite dois números inteiros (x e y):")
    fmt.Scanln(&x, &y)

    // Verificação se x é par
    
	if x%2 == 0 {
        
		// Se x for par, imprime a sequência de y números pares a partir de x
        
		fmt.Printf("Os %d números pares a partir de %d são:\n", y, x)
        
		for i := 0; i < y; i++ {
            fmt.Printf("%d ", x + 2*i)
        }
    } else {
        
		// Se x não for par, imprime uma mensagem informando que o primeiro número não é par
        fmt.Println("O primeiro número não é par.")
    }
}
