package main

import "fmt"

func main() {

	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s nice to meet you.\n", name)
	fmt.Printf("There are %d letters in your name\n", len(name))

}
