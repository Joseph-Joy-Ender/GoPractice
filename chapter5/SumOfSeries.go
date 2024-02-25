package main

import "fmt"

func seriesSum() {
	sum := 0

	fmt.Printf("%s%20s\n", "Number", "Sum")
	for number := 1; number <= 100; number++ {

		sum += number

		fmt.Printf("%4d  %20d\n", number, sum)

	}
}
