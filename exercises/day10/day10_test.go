package day10

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestGetAsteroidWithBestVisibilityExample1(t *testing.T) {
	maximumVisibleAsteroids, _ := getAsteroidWithBestVisibility([][]string{[]string{".", "#", ".", ".", "#"}, []string{".", ".", ".", ".", "."}, []string{"#", "#", "#", "#", "#"}, []string{".", ".", ".", ".", "#"}, []string{".", ".", ".", "#", "#"}})
	if maximumVisibleAsteroids != 8 {
		t.Error("Maximum visible asteroids is not 8", maximumVisibleAsteroids)
	}
}

func TestSolvePart1Input(t *testing.T) {
	maximumVisibleAsteroids, _ := getAsteroidWithBestVisibility(readInput())
	if maximumVisibleAsteroids != 303 {
		t.Error("Maximum visible asteroids is not 303", maximumVisibleAsteroids)
	}
}

func TestSolvePart2Input(t *testing.T) {
	input := readInput()

	_, asteroidCoordinates := getAsteroidWithBestVisibility(input)

	angles := getAsteroidsAtAngles(input, asteroidCoordinates.x, asteroidCoordinates.y)
	var keys []float64
	for key := range angles {
		keys = append(keys, key)
	}
	sort.Float64s(keys)

	var startIndex int
	for index, key := range keys {
		if key > -(math.Pi / 2) {
			startIndex = index - 1
			break
		}
	}

	vaporizedAsteroidCounter := 0
	for {
		currentAngle := keys[(startIndex+vaporizedAsteroidCounter)%len(keys)]
		if len(angles[currentAngle]) != 0 {
			vaporizedAsteroidCounter++

			if vaporizedAsteroidCounter == 200 {
				asteroid := angles[currentAngle][0]
				fmt.Printf("200th vaporized asteroid = (%d, %d) = %d * 100 + %d = %d\n", asteroid.x, asteroid.y, asteroid.x, asteroid.y, asteroid.x*100+asteroid.y)
				break
			}

			angles[currentAngle] = angles[currentAngle][1:]
		}
	}
}
