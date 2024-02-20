package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b int
	var c int
	var d int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	p1 := a * b
	p2 := c * d

	diferenca := fmt.Sprintf("DIFERENCA = %s", strconv.Itoa(p1-p2))

	fmt.Println(diferenca)

}
