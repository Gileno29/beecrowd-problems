package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	prod := a * b
	fmt.Println("PROD = " + strconv.Itoa(prod))
}
