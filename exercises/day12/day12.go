package day12

import (
	combinations "github.com/mxschmitt/golang-combinations"
	"strconv"
)

type Coordinate struct {
	x, y, z int
}

func getNextStep(positions []Coordinate, velocities []Coordinate) ([]Coordinate, []Coordinate) {
	// Apply gravity
	for _, combination := range combinations.All([]string{"0", "1", "2", "3"}) {
		if len(combination) == 2 {
			index0, _ := strconv.Atoi(combination[0])
			index1, _ := strconv.Atoi(combination[1])

			if positions[index0].x > positions[index1].x {
				velocities[index0].x--
				velocities[index1].x++
			} else if positions[index0].x < positions[index1].x {
				velocities[index0].x++
				velocities[index1].x--
			}

			if positions[index0].y > positions[index1].y {
				velocities[index0].y--
				velocities[index1].y++
			} else if positions[index0].y < positions[index1].y {
				velocities[index0].y++
				velocities[index1].y--
			}

			if positions[index0].z > positions[index1].z {
				velocities[index0].z--
				velocities[index1].z++
			} else if positions[index0].z < positions[index1].z {
				velocities[index0].z++
				velocities[index1].z--
			}
		}
	}

	// Apply velocity
	for positionIndex := range positions {
		positions[positionIndex].x += velocities[positionIndex].x
		positions[positionIndex].y += velocities[positionIndex].y
		positions[positionIndex].z += velocities[positionIndex].z
	}

	return positions, velocities
}
