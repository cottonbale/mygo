package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//var first string
	//var second string
	quotes := `"`
	space := ` `
	fmt.Print("quotes are" + space + quotes + "\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("hello " + text)
	fmt.Print("Enter data: ")
	text2, _ := reader.ReadString('\n')
	fmt.Println("your data " + text2)
	/*
		fmt.Printf("What is the first ?\n")
		fmt.Fscanf(&first)
		fmt.Printf("Who said it?\n")
		fmt.Fscanf(&second)
		fmt.Println(quotes + first + quotes + space + "says" + space + second)
	*/
	//fmt.Printf("%s, says, \"%s\"\n", first, second)

}
