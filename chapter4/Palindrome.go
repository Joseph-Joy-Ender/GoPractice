package main

import "strings"

func palindrome(str string) bool {
	lower := strings.ToLower(str)
	if len(lower) <= 1 {
		return true
	}

	for i := 0; i < len(lower)/2; i++ {
		if lower[i] != lower[len(lower)-i-1] {
			return false
		}
	}
	return true
}
