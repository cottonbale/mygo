package main

import "fmt"

func main() {

	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Hello ", name, " nice to meet you.\n")
	fmt.Println("length of your name is:", len(name))
	fmt.Printf("hello %s so nice to meet you.\n", name)

}
