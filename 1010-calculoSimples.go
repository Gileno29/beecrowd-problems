package main

import (
	"bufio"
	"fmt"
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
	var first string
	var second string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	first = scanner.Text()
	scanner.Scan()
	second = scanner.Text()

	unit1 := strings.Split(first, " ")[1]
	unit2 := strings.Split(second, " ")[1]
	value1 := strings.Split(first, " ")[2]
	value2 := strings.Split(second, " ")[2]

	amount1 := toFloat(unit1) * toFloat(value1)
	amount2 := toFloat(unit2) * toFloat(value2)

	result := fmt.Sprintf("VALOR A PAGAR: R$ %.2f", amount1+amount2)

	fmt.Println(result)

}
