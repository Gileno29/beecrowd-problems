package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64

	fmt.Scan(&a)

	fmt.Scan(&b)

	n1 := a * 3.5
	n2 := b * 7.5

	media := (n1 + n2) / 11

	result := fmt.Sprintf("MEDIA = %.5f", media)

	fmt.Println(result)

}
