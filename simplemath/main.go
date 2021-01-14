package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 string
	var num2 string
	//var int1 int
	//var int2 int

	fmt.Print("Enter number\n")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number\n")
	fmt.Scanln(&num2)

	int1, _ := strconv.Atoi(num1)
	int2, _ := strconv.Atoi(num2)

	total := int1 + int2
	subtracted := int1 - int2
	product := int1 * int2
	divided := int1 / int2

	fmt.Printf("answers are %d %d %d %d \n", total, subtracted, product, divided)
	fmt.Printf("%d + %d = %d\n", int1, int2, total)
	fmt.Printf("%d - %d = %d\n", int1, int2, subtracted)
	fmt.Printf("%d * %d = %d\n", int1, int2, product)
	fmt.Printf("%d / %d = %d\n", int1, int2, divided)

	//fmt.Printf("%s + %s = %s\n", strconv.Itoa(int1), strconv.Itoa(int2), strconv.Itoa(total))
	//fmt.Println("hello " + strconv.Itoa(int1))
	//fmt.Println("hello " + int1 + "\n")
}
