package day7

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

func runProgram(program []int, programCounter int, inputChannel chan string, outputChannel chan string) []int {
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
			message := <-inputChannel

			value, _ := strconv.Atoi(strings.Trim(message, "\n"))

			program[program[programCounter+1]] = value

			programCounter += 2
		}

		if opcodeString[3:] == "04" {
			value := fmt.Sprintln(getValueWithMode(program, program[programCounter+1], string(opcodeString[2])))
			outputChannel <- value

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

func getPermutationsWithHeapsAlgorithm(k int, values []int, acc *[][]int) {
	if k == 1 {
		*acc = append(*acc, cloneArray(values))
	} else {
		getPermutationsWithHeapsAlgorithm(k-1, values, acc)

		for index := 0; index < k-1; index++ {
			var aux int = values[k-1]
			if k%2 == 0 {
				values[k-1] = values[index]
				values[index] = aux
			} else {
				values[k-1] = values[0]
				values[0] = aux
			}
			getPermutationsWithHeapsAlgorithm(k-1, values, acc)
		}
	}
}

func getMaximumAmplifiedSignal(program []int) int {
	permutations := [][]int{}
	getPermutationsWithHeapsAlgorithm(5, []int{0, 1, 2, 3, 4}, &permutations)

	maxInput := 0
	for _, phaseSettings := range permutations {
		input := "0"

		for _, phase := range phaseSettings {
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			go runProgram(cloneArray(program), 0, inputChannel, outputChannel)

			inputChannel <- strconv.Itoa(phase)
			inputChannel <- input
			message := <-outputChannel
			input = strings.Trim(message, "\n")
		}

		x, _ := strconv.Atoi(input)
		if x > maxInput {
			maxInput = x
		}
	}

	return maxInput
}

func getMaximumAmplifiedSignalWithFeedbackLoop(program []int) int {
	permutations := [][]int{}
	getPermutationsWithHeapsAlgorithm(5, []int{5, 6, 7, 8, 9}, &permutations)

	maxInput := 0
	for _, phaseSettings := range permutations {
		input := "0"

		var inputChannels []chan string

		for range phaseSettings {
			inputChannel := make(chan string)
			inputChannels = append(inputChannels, inputChannel)
		}
		inputChannels = append(inputChannels, make(chan string))

		for index, phase := range phaseSettings {
			go runProgram(cloneArray(program), 0, inputChannels[index], inputChannels[index+1])
			inputChannels[index] <- strconv.Itoa(phase)
		}

		inputChannels[0] <- input

	loop:
		for {
			input = <-inputChannels[len(phaseSettings)]
			select {
			case inputChannels[0] <- strings.Trim(input, "\n"):
			default:
				break loop
			}
		}

		x, _ := strconv.Atoi(strings.Trim(input, "\n"))
		if x > maxInput {
			maxInput = x
		}
	}

	return maxInput
}

func cloneArray(a []int) []int {
	return append(a[:0:0], a...)
}
