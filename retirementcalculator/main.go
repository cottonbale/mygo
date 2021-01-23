package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	p := fmt.Println

	//ask for retirement age
	p("what is your age?")
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())

	//ask for retirement age
	p("what age do you wish to retire?")
	scanner.Scan()
	retirement, _ := strconv.Atoi(scanner.Text())

	//subtract age from retirement age to give remaining years
	yearsLeft := (retirement - age)
	p("you have " + strconv.Itoa(yearsLeft) + " years left")

	//get current date
	now := time.Now()
	nowYear := now.Year()
	//fmt.Printf("now is %d\n", nowYear)

	then := now.AddDate(yearsLeft, 0, 0)
	//p("then is " + strconv.(then))

	retYear := then.Year()

	fmt.Printf("It's %d, so you can retire in %d\n", nowYear, retYear)

	//add remaining years to current date and print

}
