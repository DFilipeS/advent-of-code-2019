package day10

import (
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type Coordinate struct {
	x, y     int
	distance float64
}

func getPolarCoordinates(x float64, y float64) (float64, float64) {
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	angle := math.Atan2(y, x)

	return r, angle
}

func getAsteroidWithBestVisibility(asteroids [][]string) (int, Coordinate) {
	var maximumVisible int
	var maximumVisibleAsteroid Coordinate
	for y := 0; y < len(asteroids); y++ {
		for x := 0; x < len(asteroids[y]); x++ {
			if asteroids[y][x] == "#" {
				visibleAsteroids := len(getAsteroidsAtAngles(asteroids, x, y))
				if visibleAsteroids > maximumVisible {
					maximumVisible = visibleAsteroids
					maximumVisibleAsteroid = Coordinate{x, y, 0}
				}
			}
		}
	}
	return maximumVisible, maximumVisibleAsteroid
}

func getAsteroidsAtAngles(asteroids [][]string, centerX int, centerY int) map[float64][]Coordinate {
	var uniqueAngles map[float64][]Coordinate = make(map[float64][]Coordinate)
	for y := 0; y < len(asteroids); y++ {
		for x := 0; x < len(asteroids[y]); x++ {
			if x == centerX && y == centerY {
				continue
			}

			if asteroids[y][x] == "#" {
				_, angle := getPolarCoordinates(float64(x-centerX), float64(y-centerY))

				uniqueAngles[angle] = append(uniqueAngles[angle], Coordinate{x, y, getDistanceBetweenPoints(centerX, centerY, x-centerX, y-centerY)})
				sort.Slice(uniqueAngles[angle], func(i, j int) bool {
					return uniqueAngles[angle][i].distance < uniqueAngles[angle][j].distance
				})
			}
		}
	}
	// fmt.Printf("(%d, %d) %v\n", centerX, centerY, uniqueAngles)
	return uniqueAngles
}

func getDistanceBetweenPoints(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2)-float64(x1), 2) + math.Pow(float64(y2)-float64(y1), 2))
}

func readInput() [][]string {
	inputBytes, _ := ioutil.ReadFile("./input.txt")

	rowsList := strings.Split(string(inputBytes), "\n")

	var asteroids [][]string
	for _, positionsList := range rowsList {
		asteroids = append(asteroids, strings.Split(positionsList, ""))
	}

	return asteroids
}
