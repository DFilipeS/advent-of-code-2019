package day5

import (
	"bufio"
	"fmt"
	"io"
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

func runProgram(program []int, programCounter int, writter io.Writer, reader io.Reader) []int {
	for {
		opcodeString := fmt.Sprintf("%05d", program[programCounter])

		if opcodeString[3:] == "01" {
			firstValue := getValueWithMode(program, program[programCounter+1], string(opcodeString[2]))
			secondValue := getValueWithMode(program, program[programCounter+2], string(opcodeString[1]))

			program[program[programCounter+3]] = firstValue + secondValue

			programCounter += 4
		}

		if opcodeString[3:] == "02" {
			firstValue := getValueWithMode(program, program[programCounter+1], string(opcodeString[2]))
			secondValue := getValueWithMode(program, program[programCounter+2], string(opcodeString[1]))

			program[program[programCounter+3]] = firstValue * secondValue

			programCounter += 4
		}

		if opcodeString[3:] == "03" {
			reader := bufio.NewReader(reader)
			text, _ := reader.ReadString('\n')
			value, _ := strconv.Atoi(strings.Trim(text, "\n"))

			program[program[programCounter+1]] = value

			programCounter += 2
		}

		if opcodeString[3:] == "04" {
			fmt.Fprintln(writter, getValueWithMode(program, program[programCounter+1], string(opcodeString[2])))

			programCounter += 2
		}

		if opcodeString[3:] == "05" {
			firstParameter := getValueWithMode(program, program[programCounter+1], string(opcodeString[2]))

			if firstParameter != 0 {
				programCounter = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]))
			} else {
				programCounter += 3
			}
		}

		if opcodeString[3:] == "06" {
			firstParameter := getValueWithMode(program, program[programCounter+1], string(opcodeString[2]))

			if firstParameter == 0 {
				programCounter = getValueWithMode(program, program[programCounter+2], string(opcodeString[1]))
			} else {
				programCounter += 3
			}
		}

		if opcodeString[3:] == "07" {
			firstParameter := getValueWithMode(program, program[programCounter+1], string(opcodeString[2]))
			secondParameter := getValueWithMode(program, program[programCounter+2], string(opcodeString[1]))

			if firstParameter < secondParameter {
				program[program[programCounter+3]] = 1
			} else {
				program[program[programCounter+3]] = 0
			}

			programCounter += 4
		}

		if opcodeString[3:] == "08" {
			firstParameter := getValueWithMode(program, program[programCounter+1], string(opcodeString[2]))
			secondParameter := getValueWithMode(program, program[programCounter+2], string(opcodeString[1]))

			if firstParameter == secondParameter {
				program[program[programCounter+3]] = 1
			} else {
				program[program[programCounter+3]] = 0
			}

			programCounter += 4
		}

		if opcodeString[3:] == "99" {
			return program
		}
	}
}

func getValueWithMode(program []int, position int, mode string) int {
	if string(mode) == "0" {
		return program[position]
	} else if string(mode) == "1" {
		return position
	}

	return -1
}
