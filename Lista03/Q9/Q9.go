package main

import (
	"fmt"
	"math"
)

// Struct para representar um ponto tridimensional
type Ponto struct {
	x, y, z float64
}

// Função para calcular a distância entre dois pontos tridimensionais
func distancia(a, b Ponto) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	// Ler o número de pontos
	var n int
	fmt.Println("Digite o número de pontos:")
	fmt.Scan(&n)

	// Verificar se o número de pontos é válido
	if n < 2 {
		fmt.Println("É necessário fornecer pelo menos 2 pontos para calcular a distância.")
		return
	}

	// Ler as coordenadas dos pontos
	pontos := make([]Ponto, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Digite as coordenadas do ponto %c (x y z): ", 'A'+i)
		fmt.Scan(&pontos[i].x, &pontos[i].y, &pontos[i].z)
	}

	// Calcular as distâncias entre pontos consecutivos
	for i := 0; i < n-1; i++ {
		distanciaAB := distancia(pontos[i], pontos[i+1])
		fmt.Printf("Distância entre %c e %c: %.2f\n", 'A'+i, 'A'+i+1, distanciaAB)
	}
}
