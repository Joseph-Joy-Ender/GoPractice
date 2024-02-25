package main

import "fmt"

func largestNumber() {
	fmt.Println("Enter ten integers: ")
	var largestNumber int
	fmt.Scanln(&largestNumber)

	for i := 1; i <= 10; i++ {
		fmt.Println("Enter ten integers: ")
		var secondNumber int
		fmt.Scanln(&secondNumber)

		if secondNumber > largestNumber {
			largestNumber = secondNumber
		}
	}
	fmt.Println("Largest number is ", largestNumber)
}
