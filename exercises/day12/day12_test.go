package day12

import (
	"fmt"
	"math"
	"testing"
)

func TestGetNextStep(t *testing.T) {
	positions := []Coordinate{Coordinate{-1, 0, 2}, Coordinate{2, -10, -7}, Coordinate{4, -8, 8}, Coordinate{3, 5, -1}}
	velocities := []Coordinate{Coordinate{0, 0, 0}, Coordinate{0, 0, 0}, Coordinate{0, 0, 0}, Coordinate{0, 0, 0}}

	var potentialEnergy float64
	var kinecticEnergy float64
	var totalEnergy float64

	for i := 0; i < 10; i++ {
		positions, velocities = getNextStep(positions, velocities)

		potentialEnergy = 0
		kinecticEnergy = 0
		totalEnergy = 0
		for index := range positions {
			potentialEnergy = math.Abs(float64(positions[index].x)) + math.Abs(float64(positions[index].y)) + math.Abs(float64(positions[index].z))
			kinecticEnergy = math.Abs(float64(velocities[index].x)) + math.Abs(float64(velocities[index].y)) + math.Abs(float64(velocities[index].z))
			totalEnergy += (potentialEnergy * kinecticEnergy)
		}
		fmt.Printf("pos=%v, vel=%v, total energy = %f\n", positions, velocities, totalEnergy)
	}
}

func TestSolvePart1Input(t *testing.T) {
	positions := []Coordinate{Coordinate{-7, -1, 6}, Coordinate{6, -9, -9}, Coordinate{-12, 2, -7}, Coordinate{4, -17, -12}}
	velocities := []Coordinate{Coordinate{0, 0, 0}, Coordinate{0, 0, 0}, Coordinate{0, 0, 0}, Coordinate{0, 0, 0}}

	var potentialEnergy float64
	var kinecticEnergy float64
	var totalEnergy float64

	for i := 0; i < 1000; i++ {
		positions, velocities = getNextStep(positions, velocities)

		potentialEnergy = 0
		kinecticEnergy = 0
		totalEnergy = 0
		for index := range positions {
			potentialEnergy = math.Abs(float64(positions[index].x)) + math.Abs(float64(positions[index].y)) + math.Abs(float64(positions[index].z))
			kinecticEnergy = math.Abs(float64(velocities[index].x)) + math.Abs(float64(velocities[index].y)) + math.Abs(float64(velocities[index].z))
			totalEnergy += (potentialEnergy * kinecticEnergy)
		}
		fmt.Printf("pos=%v, vel=%v, total energy = %f\n", positions, velocities, totalEnergy)
	}
}

func TestSolvePart2Input(t *testing.T) {
	startingPositions := []Coordinate{Coordinate{-7, -1, 6}, Coordinate{6, -9, -9}, Coordinate{-12, 2, -7}, Coordinate{4, -17, -12}}
	positions := []Coordinate{Coordinate{-7, -1, 6}, Coordinate{6, -9, -9}, Coordinate{-12, 2, -7}, Coordinate{4, -17, -12}}
	velocities := []Coordinate{Coordinate{0, 0, 0}, Coordinate{0, 0, 0}, Coordinate{0, 0, 0}, Coordinate{0, 0, 0}}

	var steps int
	var periods map[string]int = make(map[string]int)

	for {
		positions, velocities = getNextStep(positions, velocities)
		steps++

		if periods["x"] == 0 && positions[0].x == startingPositions[0].x && positions[1].x == startingPositions[1].x && positions[2].x == startingPositions[2].x && positions[3].x == startingPositions[3].x && velocities[0].x == 0 && velocities[1].x == 0 && velocities[2].x == 0 && velocities[3].x == 0 {
			periods["x"] = steps
		}

		if periods["y"] == 0 && positions[0].y == startingPositions[0].y && positions[1].y == startingPositions[1].y && positions[2].y == startingPositions[2].y && positions[3].y == startingPositions[3].y && velocities[0].y == 0 && velocities[1].y == 0 && velocities[2].y == 0 && velocities[3].y == 0 {
			periods["y"] = steps
		}

		if periods["z"] == 0 && positions[0].z == startingPositions[0].z && positions[1].z == startingPositions[1].z && positions[2].z == startingPositions[2].z && positions[3].z == startingPositions[3].z && velocities[0].z == 0 && velocities[1].z == 0 && velocities[2].z == 0 && velocities[3].z == 0 {
			periods["z"] = steps
		}

		if len(periods) == 3 {
			break
		}
	}

	if LCM(periods["x"], periods["y"], periods["z"]) != 452582583272768 {
		t.Error("Steps until repeated state is not 452582583272768", LCM(periods["x"], periods["y"], periods["z"]))
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
