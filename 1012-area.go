package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func toFloat(number string) float64 {
	n, err := strconv.ParseFloat(number, 8)

	if err != nil {
		return 0.00
	}
	return n

}
func main() {
	const PI = 3.14159
	var in string
	var a float64
	var b float64
	var c float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in = scanner.Text()
	a = toFloat(strings.Split(in, " ")[0])
	b = toFloat(strings.Split(in, " ")[1])
	c = toFloat(strings.Split(in, " ")[2])

	triangulo := fmt.Sprintf("TRIANGULO: %.3f", (a*c)/2)
	circulo := fmt.Sprintf("CIRCULO: %.3f", PI*(math.Pow(c, 2)))
	trapezio := fmt.Sprintf("TRAPEZIO: %.3f", (a+b)*(c/2))
	quadrado := fmt.Sprintf("QUADRADO: %.3f", math.Pow(b, 2))
	retangulo := fmt.Sprintf("RETANGULO: %.3f", a*b)

	fmt.Println(triangulo)
	fmt.Println(circulo)
	fmt.Println(trapezio)
	fmt.Println(quadrado)
	fmt.Println(retangulo)

}
