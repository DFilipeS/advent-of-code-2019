package day8

import (
	"fmt"
	"testing"
)

func TestSolvePart1Input(t *testing.T) {
	layer := findLayerWithFewestZeroes(splitLayers(readInput(), 25, 6))

	totalOnes := 0
	totalTwos := 0
	for _, pixel := range layer {
		if pixel == 1 {
			totalOnes++
		}
		if pixel == 2 {
			totalTwos++
		}
	}
	fmt.Printf("The answer is: %d * %d = %d\n", totalOnes, totalTwos, totalOnes*totalTwos)
}

func TestSolvePart2Input(t *testing.T) {
	width := 25
	height := 6
	decodeImage(splitLayers(readInput(), width, height), width, height)
}
