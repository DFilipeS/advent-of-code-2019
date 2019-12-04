package day4

import (
	"bytes"
)

// Part 1

func validatePassword(password string) bool {
	runes := bytes.Runes([]byte(password))
	hasRepeated := false
	isNonDecreasing := true

	for index, rune := range runes {
		if index < len(runes)-1 {
			if rune == runes[index+1] {
				hasRepeated = true
			}

			if rune > runes[index+1] {
				isNonDecreasing = false
			}
		}
	}

	return hasRepeated && isNonDecreasing
}

// Part 2

func validatePasswordExtended(password string) bool {
	runes := bytes.Runes([]byte(password))
	hasRepeated := false
	digitsCount := map[rune]int{}
	isNonDecreasing := true

	for index, r := range runes {
		digitsCount[r]++

		if index < len(runes)-1 {
			if r == runes[index+1] {
				hasRepeated = true
			}

			if r > runes[index+1] {
				isNonDecreasing = false
			}
		}
	}

	isExtendedValid := false
	for _, value := range digitsCount {
		if value == 2 {
			isExtendedValid = true
		}
	}

	return hasRepeated && isNonDecreasing && isExtendedValid
}
