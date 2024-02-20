package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	n1 := a * 2
	n2 := b * 3
	n3 := c * 5

	media := fmt.Sprintf("MEDIA = %.1f", (n1+n2+n3)/10)
	fmt.Println(media)

}
