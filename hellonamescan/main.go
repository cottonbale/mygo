package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter name: ")
	scanner.Scan()
	text := scanner.Text()

	fmt.Printf("Hello %s, your name is %d characters long\n", text, len(text))

}
