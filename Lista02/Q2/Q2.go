package main

import (
	"fmt"
)

var popA, popB, anos int = 0, 0, 0
const (
	txA = 1.03;
	txB = 1.015;
)

func main() {

	fmt.Printf("População A: ")
	fmt.Scan(&popA)

	fmt.Printf("População B: ")
	fmt.Scan(&popB)

	for popA < popB {

		popA = int(float64(popA) * txA)
		popB = int(float64(popB) * txB)

		anos++
	}

	fmt.Print("Anos:", anos)

}