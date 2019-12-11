package day9

import (
	"fmt"
	"strings"
	"testing"
)

func TestBoostExample1(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	go runProgram(buildProgramMap([]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}), 0, inputChannel, outputChannel)

	for {
		output := <-outputChannel
		fmt.Printf("%s ", strings.Trim(output, "\n"))
		if output == "99\n" {
			fmt.Println()
			break
		}
	}
}

func TestBoostExample2(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	go runProgram(buildProgramMap([]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}), 0, inputChannel, outputChannel)

	output := <-outputChannel
	if output != "1219070632396864\n" {
		t.Error("Control value is not 1219070632396864", output)
	}
}

func TestBoostExample3(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	go runProgram(buildProgramMap([]int{104, 1125899906842624, 99}), 0, inputChannel, outputChannel)

	output := <-outputChannel
	if output != "1125899906842624\n" {
		t.Error("Control value is not 1125899906842624", output)
	}
}

func TestSolvePart1Input(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	go runProgram(buildProgramMap(readInput()), 0, inputChannel, outputChannel)

	inputChannel <- "1\n"
	output := <-outputChannel
	if output != "3100786347\n" {
		t.Error("Keycode is not 3100786347", output)
	}
}

func TestSolvePart2Input(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	go runProgram(buildProgramMap(readInput()), 0, inputChannel, outputChannel)

	inputChannel <- "2\n"
	output := <-outputChannel
	if output != "87023\n" {
		t.Error("Coordinates are not 87023", output)
	}
}
