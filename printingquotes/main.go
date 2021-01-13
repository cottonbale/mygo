package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	quotes := `"`
	space := ` `

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the quote ? ")
	scanner.Scan()
	first := scanner.Text()

	fmt.Print("Who said it ? ")
	scanner.Scan()
	second := scanner.Text()

	fmt.Println(quotes + first + quotes + space + "says" + space + second)

}
