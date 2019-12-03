package day2

import (
	"fmt"
	"testing"
)

func TestRunProgram(t *testing.T) {
	exampleProgram1 := runProgram([]int{1, 0, 0, 0, 99}, 0)
	if exampleProgram1[0] != 2 {
		t.Error("runProgram of {1, 0, 0, 0, 99} is not {2, 0, 0, 0, 99}", exampleProgram1)
	}

	exampleProgram2 := runProgram([]int{2, 3, 0, 3, 99}, 0)
	if exampleProgram2[3] != 6 {
		t.Error("runProgram of {2,3,0,3,99} is not {2,3,0,6,99}", exampleProgram2)
	}

	exampleProgram3 := runProgram([]int{2, 4, 4, 5, 99, 0}, 0)
	if exampleProgram3[5] != 9801 {
		t.Error("runProgram of {2,4,4,5,99,0} is not {2,4,4,5,99,9801}", exampleProgram3)
	}
}

func TestSolvePart1Input(t *testing.T) {
	program := readInput()
	program[1] = 12
	program[2] = 2

	program = runProgram(program, 0)
	if program[0] != 3409710 {
		t.Error("Position 0 is not 3409710", program[0])
	}
}

func TestSolvePart2Input(t *testing.T) {
	noun, verb := findNounAndVerb(19690720)

	fmt.Printf("noun = %d; verb = %d; 100 * noun + verb = %d", noun, verb, 100*noun+verb)
}
