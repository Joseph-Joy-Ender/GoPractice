package main

import "fmt"

func collectInput() int {
	fmt.Println("Enter a number between 1 and 30: ")
	var number int
	fmt.Scanln(&number)
	return number
}

func printAsterisk(input int) {
	for i := 0; i < input; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func printing() {
	firstInput := collectInput()
	secondInput := collectInput()
	thirdInput := collectInput()
	fourthInput := collectInput()
	fifthInput := collectInput()

	fmt.Println()

	printAsterisk(firstInput)
	printAsterisk(secondInput)
	printAsterisk(thirdInput)
	printAsterisk(fourthInput)
	printAsterisk(fifthInput)
}
