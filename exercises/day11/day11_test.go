package day11

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestSolvePart1Input(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	go runProgram(buildProgramMap(readInput()), 0, inputChannel, outputChannel)

	var currentPosition = Coordinate{0, 0}
	var orientation string = "UP"
	var paintedPanels map[Coordinate]string = make(map[Coordinate]string)
	paintedPanels[currentPosition] = "0"

	for {
		message := strings.Trim(<-outputChannel, "\n")

		if message == "EOP" {
			break
		}

		inputChannel <- fmt.Sprintf("%s\n", paintedPanels[currentPosition])
		color := strings.Trim(<-outputChannel, "\n")

		turnDirection := strings.Trim(<-outputChannel, "\n")

		paintedPanels[currentPosition] = color

		orientation = getNextOrientation(orientation, turnDirection)
		currentPosition = getNextPosition(currentPosition, orientation)
	}

	fmt.Println("Total unique painted panels:", len(paintedPanels))
}

func TestSolvePart2Input(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	go runProgram(buildProgramMap(readInput()), 0, inputChannel, outputChannel)

	var currentPosition = Coordinate{0, 0}
	var orientation string = "UP"
	var paintedPanels map[Coordinate]string = make(map[Coordinate]string)
	paintedPanels[currentPosition] = "1"

	for {
		message := strings.Trim(<-outputChannel, "\n")

		if message == "EOP" {
			break
		}

		inputChannel <- fmt.Sprintf("%s\n", paintedPanels[currentPosition])
		color := strings.Trim(<-outputChannel, "\n")

		turnDirection := strings.Trim(<-outputChannel, "\n")

		paintedPanels[currentPosition] = color

		orientation = getNextOrientation(orientation, turnDirection)
		currentPosition = getNextPosition(currentPosition, orientation)
	}

	var maxX, maxY int
	var minX, minY int = math.MaxInt64, math.MaxInt64

	for key := range paintedPanels {
		if key.x > maxX {
			maxX = key.x
		}
		if key.y > maxY {
			maxY = key.y
		}

		if key.x < minX {
			minX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if paintedPanels[Coordinate{x, y}] == "1" {
				fmt.Print("1")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
