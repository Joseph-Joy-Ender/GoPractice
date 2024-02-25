package main

import "fmt"

func twoLargestNumber() {
	fmt.Println("Enter ten integers: ")
	var number int
	fmt.Scanln(&number)

	firstLargestNumber := number
	secondLargestNUmber := number

	for count := 1; count <= 10; count++ {
		fmt.Println("Enter ten integers: ")
		var secondNumber int
		fmt.Scanln(&secondNumber)

		if secondNumber > firstLargestNumber {
			secondLargestNUmber = firstLargestNumber
			firstLargestNumber = secondNumber
		}
	}

	fmt.Println("First largest number is: ", firstLargestNumber)
	fmt.Println("Second largest number is: ", secondLargestNUmber)
}
