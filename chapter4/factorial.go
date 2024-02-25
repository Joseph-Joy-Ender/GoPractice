package main

func factorial(number uint) uint {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
