package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = 3.14159

	var raio float64

	fmt.Scan(&raio)
	area := pi * (math.Pow(raio, 2))
	/**
	 * Escreva a sua solução aqui
	 * Code your solution here
	 * Escriba su solución aquí
	 */
	result := fmt.Sprintf("A=%.4f", area)
	fmt.Println(result)

}
