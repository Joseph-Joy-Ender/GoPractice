package main

import "fmt"

func studentGrade() {
	aCounter := 0
	bCounter := 0
	cCounter := 0
	dCounter := 0

	for count := 1; count <= 5; count++ {
		fmt.Println("Enter your name: ")
		var name string
		fmt.Scanln(&name)

		fmt.Println("Enter your grade: ")
		var grade string
		fmt.Scanln(&grade)

		switch grade {
		case "A":
			aCounter++
		case "B":
			bCounter++
		case "C":
			cCounter++
		case "D":
			dCounter++

		}
		fmt.Println("Name of student is ", name, "\n", "Grade is ", grade)

	}

	fmt.Printf("\n%s\n%s%d\n%s%d\n%s%d\n%s%d",
		"Result of student grade:",
		"A: ", aCounter,
		"B: ", bCounter,
		"C: ", cCounter,
		"D: ", dCounter,
	)
}
