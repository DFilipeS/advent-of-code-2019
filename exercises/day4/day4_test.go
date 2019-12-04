package day4

import (
	"fmt"
	"strconv"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	if !validatePassword("111111") {
		t.Error("Password 111111 should be valid")
	}

	if validatePassword("223450") {
		t.Error("Password 223450 should not be valid")
	}

	if validatePassword("123789") {
		t.Error("Password 123789 should not be valid")
	}
}

func TestValidatePasswordExtended(t *testing.T) {
	if !validatePasswordExtended("112233") {
		t.Error("Password 112233 should be valid")
	}

	if validatePasswordExtended("123444") {
		t.Error("Password 123444 should not be valid")
	}

	if !validatePasswordExtended("111122") {
		t.Error("Password 111122 should be valid")
	}
}

func TestSolvePart1Input(t *testing.T) {
	validPassowordsCounter := 0

	for password := 172930; password < 683082; password++ {
		if validatePassword(strconv.Itoa(password)) {
			validPassowordsCounter++
		}
	}

	fmt.Printf("Total valid passwords = %d\n", validPassowordsCounter)
}

func TestSolvePart2Input(t *testing.T) {
	validPassowordsCounter := 0

	for password := 172930; password < 683082; password++ {
		if validatePasswordExtended(strconv.Itoa(password)) {
			validPassowordsCounter++
		}
	}

	fmt.Printf("Total valid passwords (extended) = %d\n", validPassowordsCounter)
}
