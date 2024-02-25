package main

import "fmt"

func tabularOutput() {
	fmt.Println("N               N1             N2              N3               N4")

	for i := 1; i <= 5; i++ {
		fmt.Print(i)

		currentValue := i

		for j := 0; j < 4; j++ {
			fmt.Print("\t\t")
			currentValue = currentValue * i
			fmt.Print(currentValue)

		}
		fmt.Println()
	}
}
