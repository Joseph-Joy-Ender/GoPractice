package main

import "fmt"

func specifiedSum() {
	fmt.Println("Enter a number: ")
	var initialNumber int
	fmt.Scanln(&initialNumber)

	total := 0

	for total < initialNumber {
		fmt.Println("Enter numbers: ")
		var number int
		fmt.Scanln(&number)

		total += number
	}
	fmt.Println(total)

}
