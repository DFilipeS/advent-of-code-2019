package day3

import (
	"bytes"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Coordinate - Struct to hold a simple coordinate in a 2D plane
type Coordinate struct {
	X, Y int
}

func readInput() [][]string {
	inputBytes, _ := ioutil.ReadFile("./input.txt")
	wiresList := strings.Split(string(inputBytes), "\n")
	var wiresStepsList [][]string

	for _, wireString := range wiresList {
		stepsList := strings.Split(wireString, ",")
		wiresStepsList = append(wiresStepsList, stepsList)
	}

	return wiresStepsList
}

// Part 1

func calculateMinimumDistance(wiresStepsList [][]string) float64 {
	intersectionsMapSet, _ := createIntersectionsMapSet(wiresStepsList)

	var minDistance float64 = math.MaxFloat64
	for key := range intersectionsMapSet {
		if key.X != 0 && key.Y != 0 {
			distance := math.Abs(float64(key.X)) + math.Abs(float64(key.Y))

			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}

func createIntersectionsMapSet(wiresStepsList [][]string) (map[Coordinate]bool, map[Coordinate]map[int]int) {
	gridMap := map[Coordinate]map[int]int{}
	intersectionsMapSet := map[Coordinate]bool{}

	for wireIndex, stepsList := range wiresStepsList {
		currentCoordinate := Coordinate{0, 0}
		totalSteps := 0

		for _, stepString := range stepsList {
			if stepString != "" {
				runes := bytes.Runes([]byte(stepString))
				direction := string(runes[0:1])
				steps, _ := strconv.Atoi(string(runes[1:]))

				for index := 0; index < steps; index++ {
					if direction == "U" {
						currentCoordinate.Y++
					} else if direction == "D" {
						currentCoordinate.Y--
					} else if direction == "L" {
						currentCoordinate.X--
					} else if direction == "R" {
						currentCoordinate.X++
					}

					if gridMap[currentCoordinate] == nil {
						gridMap[currentCoordinate] = make(map[int]int)
					}

					gridMap[currentCoordinate][wireIndex] = totalSteps + 1
					totalSteps++

					if len(gridMap[currentCoordinate]) > 1 {
						intersectionsMapSet[currentCoordinate] = true
					}
				}
			}
		}
	}

	return intersectionsMapSet, gridMap
}

// Part 2

func calculateMinimumSteps(wiresStepsList [][]string) int {
	intersectionsMapSet, gridMap := createIntersectionsMapSet(wiresStepsList)

	var minSteps int = math.MaxInt32
	for coordinate := range intersectionsMapSet {
		var totalSteps int
		for _, steps := range gridMap[coordinate] {
			totalSteps += steps
		}

		if totalSteps < minSteps {
			minSteps = totalSteps
		}
	}

	return minSteps
}
