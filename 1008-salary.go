package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var hours int
	var salaryPerHour float64

	fmt.Scan(&number)
	fmt.Scan(&hours)
	fmt.Scan(&salaryPerHour)

	fmt.Println("NUMBER = " + strconv.Itoa(number))
	salary := fmt.Sprintf("SALARY = U$ %.2f", salaryPerHour*float64(hours))

	fmt.Println(salary)

}
