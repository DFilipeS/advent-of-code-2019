package day5

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunProgram(t *testing.T) {
	var writterExample1 *bytes.Buffer = bytes.NewBufferString("")
	var readerExample1 *strings.Reader = strings.NewReader("")

	exampleProgram1 := runProgram([]int{1002, 4, 3, 4, 33}, 0, writterExample1, readerExample1)
	if exampleProgram1[4] != 99 {
		t.Error("runProgram of {1002,4,3,4,33} is not {1002,4,3,4,99}", exampleProgram1)
	}

	var writterExample2 *bytes.Buffer = bytes.NewBufferString("")
	var readerExample2 *strings.Reader = strings.NewReader("")
	exampleProgram2 := runProgram([]int{1101, 100, -1, 4, 0}, 0, writterExample2, readerExample2)
	if exampleProgram2[4] != 99 {
		t.Error("runProgram of {1101,100,-1,4,99} is not {1101,100,-1,4,99}", exampleProgram2)
	}

	var writterExample3 *bytes.Buffer = bytes.NewBufferString("")
	var readerExample3 *strings.Reader = strings.NewReader("")
	exampleProgram3 := runProgram([]int{1, 0, 0, 0, 99}, 0, writterExample3, readerExample3)
	if exampleProgram3[0] != 2 {
		t.Error("runProgram of {1, 0, 0, 0, 99} is not {2, 0, 0, 0, 99}", exampleProgram3)
	}

	var writterExample4 *bytes.Buffer = bytes.NewBufferString("")
	var readerExample4 *strings.Reader = strings.NewReader("")
	exampleProgram4 := runProgram([]int{2, 3, 0, 3, 99}, 0, writterExample4, readerExample4)
	if exampleProgram4[3] != 6 {
		t.Error("runProgram of {2,3,0,3,99} is not {2,3,0,6,99}", exampleProgram4)
	}

	var writterExample5 *bytes.Buffer = bytes.NewBufferString("")
	var readerExample5 *strings.Reader = strings.NewReader("")
	exampleProgram5 := runProgram([]int{2, 4, 4, 5, 99, 0}, 0, writterExample5, readerExample5)
	if exampleProgram5[5] != 9801 {
		t.Error("runProgram of {2,4,4,5,99,0} is not {2,4,4,5,99,9801}", exampleProgram5)
	}
}

func TestSolvePart1Input(t *testing.T) {
	program := readInput()

	var writter *bytes.Buffer = bytes.NewBufferString("")
	var reader *strings.Reader = strings.NewReader("1\n")
	program = runProgram(program, 0, writter, reader)

	output := writter.String()

	if !strings.Contains(output, "7988899") {
		t.Error("runProgram does not contain the value 7988899")
	}
}

func TestSolvePart2Input(t *testing.T) {
	program := readInput()

	var writter *bytes.Buffer = bytes.NewBufferString("")
	var reader *strings.Reader = strings.NewReader("5\n")
	program = runProgram(program, 0, writter, reader)

	output := writter.String()
	fmt.Println(output)

	if !strings.Contains(output, "13758663") {
		t.Error("runProgram does not contain the value 13758663")
	}
}
