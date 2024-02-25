package main

import "fmt"

func extremes() {
	fmt.Println("Enter numbers: ")
	var number int
	fmt.Scanln(&number)

	maximum := number
	minimum := number

	condition := true

	fmt.Println("Press (Y) for continuation and (N) to stop the program: ")
	var answer string
	fmt.Scanln(&answer)

	if answer == "N" {
		condition = false

	}

	total := 0

	for condition {
		fmt.Println("Enter numbers: ")
		var secondNumber int
		fmt.Scanln(&secondNumber)

		if secondNumber > maximum {
			maximum = secondNumber
		}
		if secondNumber < minimum {
			minimum = secondNumber
		}

		fmt.Println("Press (Y) for continuation and (N) to stop the program: ")
		fmt.Scanln(&answer)

		if answer == "N" {
			condition = false
		}

		total = maximum + minimum
	}

	fmt.Println("Maximum number is ", maximum)
	fmt.Println("Minimum number is ", minimum)
	fmt.Println("The sum of two extremes is ", total)
}
