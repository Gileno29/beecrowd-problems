package main

import (
	"fmt"
)

func main() {
	var name string
	var salary float64
	var sales float64

	fmt.Scan(&name)
	fmt.Scan(&salary)
	fmt.Scan(&sales)

	amount := salary + (sales*15)/100

	t := fmt.Sprintf("TOTAL = R$ %.2f", amount)

	fmt.Println(t)
}
