package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a noun\n")
	scanner.Scan()
	noun := scanner.Text()

	fmt.Print("Enter a verb\n")
	scanner.Scan()
	verb := scanner.Text()

	fmt.Print("Enter an adjective\n")
	scanner.Scan()
	adjective := scanner.Text()

	fmt.Print("Enter an adverb\n")
	scanner.Scan()
	adverb := scanner.Text()

	fmt.Printf("Do you %s your %s %s %s ? That's hilarious!\n", verb, adjective, noun, adverb)

}
