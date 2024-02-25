package main

import (
	"strconv"
)

func convertToBinary(binary int) int {
	var binaryString = strconv.Itoa(binary)

	var num = 0
	var exponential int = 0

	for i := len(binaryString) - 1; i >= 0; i-- {
		var bit = rune(binaryString[i])
		if bit == '1' {
			//num += math.Pow(2, exponential)
			exponential++
		}

	}
	return int(num)

}
