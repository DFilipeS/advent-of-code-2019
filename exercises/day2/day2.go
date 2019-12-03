package day2

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func findNounAndVerb(target int) (int, int) {
	var program []int

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program = readInput()
			program[1] = noun
			program[2] = verb

			program = runProgram(program, 0)
			if program[0] == target {
				return noun, verb
			}
		}
	}

	return -1, -1
}

func readInput() []int {
	inputBytes, _ := ioutil.ReadFile("./input.txt")

	stringsList := strings.Split(string(inputBytes), ",")
	var intList = make([]int, len(stringsList))

	for index, numberString := range stringsList {
		number, _ := strconv.Atoi(numberString)
		intList[index] = number
	}

	return intList
}

func runProgram(program []int, programCounter int) []int {
	for {
		opcode := program[programCounter]

		if opcode == 1 {
			positionX := program[programCounter+1]
			positionY := program[programCounter+2]
			positionZ := program[programCounter+3]

			program[positionZ] = program[positionX] + program[positionY]
		}

		if opcode == 2 {
			positionX := program[programCounter+1]
			positionY := program[programCounter+2]
			positionZ := program[programCounter+3]

			program[positionZ] = program[positionX] * program[positionY]
		}

		if opcode == 99 {
			return program
		}

		programCounter += 4
	}
}
