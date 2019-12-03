package day1

import (
	"fmt"
	"testing"
)

func TestCalculateFuelRequirementForMass(t *testing.T) {
	if calculateFuelRequirementForMass(12) != 2 {
		t.Error("calculateFuelRequirementForMass of 12 is not 2", calculateFuelRequirementForMass(12))
	}

	if calculateFuelRequirementForMass(14) != 2 {
		t.Error("calculateFuelRequirementForMass of 14 is not 2", calculateFuelRequirementForMass(14))
	}

	if calculateFuelRequirementForMass(1969) != 654 {
		t.Error("calculateFuelRequirementForMass of 1969 is not 654", calculateFuelRequirementForMass(1969))
	}

	if calculateFuelRequirementForMass(100756) != 33583 {
		t.Error("calculateFuelRequirementForMass of 100756 is not 33583", calculateFuelRequirementForMass(100756))
	}
}

func TestCalculateRecursiveFuelRequirementForMass(t *testing.T) {
	if calculateRecursiveFuelRequirementForMass(14) != 2 {
		t.Error("calculateRecursiveFuelRequirementForMass of 14 is not 2", calculateRecursiveFuelRequirementForMass(14))
	}

	if calculateRecursiveFuelRequirementForMass(1969) != 966 {
		t.Error("calculateRecursiveFuelRequirementForMass of 1969 is not 966", calculateRecursiveFuelRequirementForMass(1969))
	}

	if calculateRecursiveFuelRequirementForMass(100756) != 50346 {
		t.Error("calculateRecursiveFuelRequirementForMass of 100756 is not 50346", calculateRecursiveFuelRequirementForMass(100756))
	}
}

func TestSolvePart1Input(t *testing.T) {
	fuelRequirement := calculateFuelRequirement()
	fmt.Printf("Fuel requirement: %d\n", fuelRequirement)

	if fuelRequirement != 3454026 {
		t.Error("calculateFuelRequirement of given input is not 3454026", fuelRequirement)
	}
}

func TestSolvePart2Input(t *testing.T) {
	fuelRequirement := calculateRecursiveFuelRequirement()
	fmt.Printf("Total fuel requirement: %d\n", fuelRequirement)

	if fuelRequirement != 5178170 {
		t.Error("calculateRecursiveFuelRequirement of given input is not 5178170", fuelRequirement)
	}
}
