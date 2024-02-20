package main

import (
	"fmt"
	"math"
)

func main() {
	const formula = (4 / 3.0) * 3.14159

	var raio float64
	fmt.Scan(&raio)
	r := formula * math.Pow(raio, 3)
	result := fmt.Sprintf("VOLUME = %.3f", r)
	fmt.Println(result)

}
