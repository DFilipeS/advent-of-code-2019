package day8

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func findLayerWithFewestZeroes(layers [][]int) []int {
	minimumZeroes := math.MaxInt32
	var minimumZeroesLayer []int
	for _, layer := range layers {
		totalZeroes := 0
		for _, pixel := range layer {
			if pixel == 0 {
				totalZeroes++
			}
		}

		if totalZeroes < minimumZeroes {
			minimumZeroes = totalZeroes
			minimumZeroesLayer = layer
		}
	}

	return minimumZeroesLayer
}

func decodeImage(layers [][]int, width int, height int) {
	image := layers[0]

	for layerIndex := 1; layerIndex < len(layers); layerIndex++ {
		for pixelIndex, pixel := range layers[layerIndex] {
			if image[pixelIndex] == 2 && pixel != 2 {
				image[pixelIndex] = pixel
			}
		}
	}

	printImage(image, width, height)
}

func printImage(image []int, width int, height int) {
	for i := 0; i < width*height; i++ {
		if image[i] == 1 {
			fmt.Print(image[i])
		} else {
			fmt.Print(" ")
		}
		if i%width == width-1 {
			fmt.Print("\n")
		}
	}
}

func readInput() []int {
	inputBytes, _ := ioutil.ReadFile("./input.txt")

	stringsList := strings.Split(string(inputBytes), "")
	var intList = make([]int, len(stringsList))

	for index, numberString := range stringsList {
		number, _ := strconv.Atoi(numberString)
		intList[index] = number
	}

	return intList
}

func splitLayers(pixels []int, width int, height int) [][]int {
	return chunkArray(pixels, width*height)
}

func chunkArray(pixels []int, chunkSize int) [][]int {
	var layers [][]int
	for index := 0; index < len(pixels); index += chunkSize {
		layers = append(layers, pixels[index:index+chunkSize])
	}
	return layers
}
