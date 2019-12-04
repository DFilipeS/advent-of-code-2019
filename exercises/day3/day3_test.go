package day3

import (
	"fmt"
	"testing"
)

func TestReadInput(t *testing.T) {
	readInput()
}

func TestCalculateMinimumDistance(t *testing.T) {
	minimumDistanceExample1 := calculateMinimumDistance([][]string{{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, {"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}})
	if minimumDistanceExample1 != 159 {
		t.Error("minimumDistance of first example is not 159", minimumDistanceExample1)
	}

	minimumDistanceExample2 := calculateMinimumDistance([][]string{{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, {"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}})
	if minimumDistanceExample2 != 135 {
		t.Error("minimumDistance of second example is not 135", minimumDistanceExample2)
	}
}

func TestSolvePart1Input(t *testing.T) {
	minimumDistance := calculateMinimumDistance(readInput())

	fmt.Printf("minimumDistance = %f\n", minimumDistance)
}

func TestCalculateMinimumSteps(t *testing.T) {
	minimumStepsExample1 := calculateMinimumSteps([][]string{{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, {"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}})
	if minimumStepsExample1 != 610 {
		t.Error("minimumSteps of first example is not 610", minimumStepsExample1)
	}

	minimumDistanceExample2 := calculateMinimumSteps([][]string{{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, {"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}})
	if minimumDistanceExample2 != 410 {
		t.Error("minimumSteps of second example is not 410", minimumDistanceExample2)
	}
}

func TestSolvePart2Input(t *testing.T) {
	minimumSteps := calculateMinimumSteps(readInput())

	fmt.Printf("minimumSteps = %d\n", minimumSteps)
}
