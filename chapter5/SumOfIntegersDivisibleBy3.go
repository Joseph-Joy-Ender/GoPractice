package main

import "fmt"

func divisibleByThree() {
	total := 0

	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			total += i
		}
	}

	fmt.Println("The sum of integers between 1 and 30 that is divisible by 3 is ", total)
}
