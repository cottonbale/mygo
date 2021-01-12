package main

import "fmt"

func main() {

	var name string

	fmt.Print("What is the input string ?\n")
	fmt.Scanln(&name)
	fmt.Printf("%s has %d characters\n", name, len(name))

}
