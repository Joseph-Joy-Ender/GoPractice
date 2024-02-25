package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
		//fmt.Println(i)
	}

	fmt.Println()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Yes")
		} else if i%3 == 0 {
			fmt.Println("No")
		} else if i%2 == 0 {
			fmt.Println("YesNo")
		} else {
			fmt.Println("What")
		}
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		default:
			fmt.Println("Unknown Number")

		}

	}

}
