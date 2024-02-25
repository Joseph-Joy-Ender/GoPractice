package main

import "fmt"

func main() {
	var num [5]int
	num[4] = 100
	fmt.Println(num)

	var nums [5]float64
	nums[0] = 86
	nums[1] = 75
	nums[2] = 89
	nums[3] = 77
	nums[4] = 98

	var total float64 = 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	fmt.Println(total / float64(len(nums)))

}
