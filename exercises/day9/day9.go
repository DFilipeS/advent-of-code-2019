package day9

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func buildProgramMap(program []int) map[int]int {
	var programMap map[int]int = make(map[int]int)
	for i, value := range program {
		programMap[i] = value
	}
	return programMap
}

func runProgram(program map[int]int, programCounter int, inputChannel chan string, outputChannel chan string) map[int]int {
	var firstParameter int
	var secondParameter int
	var relativeBase int

	for {
		opcodeString := fmt.Sprintf("%05d", program[programCounter])

		if opcodeString[3:] == "01" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)
			secondParameter, program = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]), relativeBase)

			var offset int
			if string(opcodeString[0]) == "2" {
				offset = relativeBase
			}

			program[program[programCounter+3]+offset] = firstParameter + secondParameter

			programCounter += 4
		}

		if opcodeString[3:] == "02" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)
			secondParameter, program = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]), relativeBase)

			var offset int
			if string(opcodeString[0]) == "2" {
				offset = relativeBase
			}

			program[program[programCounter+3]+offset] = firstParameter * secondParameter

			programCounter += 4
		}

		if opcodeString[3:] == "03" {
			message := <-inputChannel

			value, _ := strconv.Atoi(strings.Trim(message, "\n"))

			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)

			if string(opcodeString[2]) == "2" {
				program[program[programCounter+1]+relativeBase] = value
			} else {
				program[program[programCounter+1]] = value
			}

			programCounter += 2
		}

		if opcodeString[3:] == "04" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)
			value := fmt.Sprintln(firstParameter)
			outputChannel <- value

			programCounter += 2
		}

		if opcodeString[3:] == "05" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)

			if firstParameter != 0 {
				programCounter, program = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]), relativeBase)
			} else {
				programCounter += 3
			}
		}

		if opcodeString[3:] == "06" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)

			if firstParameter == 0 {
				programCounter, program = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]), relativeBase)
			} else {
				programCounter += 3
			}
		}

		if opcodeString[3:] == "07" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)
			secondParameter, program = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]), relativeBase)

			var offset int
			if string(opcodeString[0]) == "2" {
				offset = relativeBase
			}

			if firstParameter < secondParameter {
				program[program[programCounter+3]+offset] = 1
			} else {
				program[program[programCounter+3]+offset] = 0
			}

			programCounter += 4
		}

		if opcodeString[3:] == "08" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)
			secondParameter, program = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]), relativeBase)

			var offset int
			if string(opcodeString[0]) == "2" {
				offset = relativeBase
			}

			if firstParameter == secondParameter {
				program[program[programCounter+3]+offset] = 1
			} else {
				program[program[programCounter+3]+offset] = 0
			}

			programCounter += 4
		}

		if opcodeString[3:] == "09" {
			firstParameter, program = getValueWithMode(program, program[programCounter+1], string(opcodeString[2]), relativeBase)
			relativeBase += firstParameter

			programCounter += 2
		}

		if opcodeString[3:] == "99" {
			return program
		}
	}
}

func getValueWithMode(program map[int]int, position int, mode string, relativeBase int) (int, map[int]int) {
	if string(mode) == "0" {
		return program[position], program
	} else if string(mode) == "1" {
		return position, program
	} else if string(mode) == "2" {
		return program[position+relativeBase], program
	}

	return -1, program
}
