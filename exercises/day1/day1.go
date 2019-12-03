package day1

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readInput() []string {
	inputBytes, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(inputBytes), "\n")
}

// Part 1

func calculateFuelRequirement() int {
	lines := readInput()

	sum := 0
	for _, n := range lines {
		if n != "" {
			mass, _ := strconv.Atoi(n)
			sum = sum + calculateFuelRequirementForMass(mass)
		}
	}

	return sum
}

func calculateFuelRequirementForMass(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

// Part 2

func calculateRecursiveFuelRequirement() int {
	lines := readInput()

	sum := 0
	for _, n := range lines {
		if n != "" {
			mass, _ := strconv.Atoi(n)
			sum = sum + calculateRecursiveFuelRequirementForMass(mass)
		}
	}

	return sum
}

func calculateRecursiveFuelRequirementForMass(mass int) int {
	fuel := int(math.Floor(float64(mass)/3)) - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + calculateRecursiveFuelRequirementForMass(fuel)
}
