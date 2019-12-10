package day7

import (
	"testing"
)

func TestGetMaximumAmplifiedSignalExample1(t *testing.T) {
	signal := getMaximumAmplifiedSignal([]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0})
	if signal != 43210 {
		t.Error("Maximum amplified signal of example 1 is not 43210", signal)
	}
}

func TestGetMaximumAmplifiedSignalExample2(t *testing.T) {
	signal := getMaximumAmplifiedSignal([]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0})
	if signal != 54321 {
		t.Error("Maximum amplified signal of example 2 is not 54321", signal)
	}
}

func TestGetMaximumAmplifiedSignalExample3(t *testing.T) {
	signal := getMaximumAmplifiedSignal([]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0})
	if signal != 65210 {
		t.Error("Maximum amplified signal of example 3 is not 65210", signal)
	}
}

func TestGetMaximumAmplifiedSignalPart1(t *testing.T) {
	signal := getMaximumAmplifiedSignal(readInput())
	if signal != 422858 {
		t.Error("Maximum amplified signal of part 1 is not 422858", signal)
	}
}

func TestGetMaximumAmplifiedSignalWithFeedbackLoopExample1(t *testing.T) {
	signal := getMaximumAmplifiedSignalWithFeedbackLoop([]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5})
	if signal != 139629729 {
		t.Error("Maximum amplified signal (with feedback loop) of example 1 is not 139629729", signal)
	}
}

func TestGetMaximumAmplifiedSignalWithFeedbackLoopExample2(t *testing.T) {
	signal := getMaximumAmplifiedSignalWithFeedbackLoop([]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10})
	if signal != 18216 {
		t.Error("Maximum amplified signal (with feedback loop) of example 2 is not 18216", signal)
	}
}

func TestGetMaximumAmplifiedSignalPart2(t *testing.T) {
	signal := getMaximumAmplifiedSignalWithFeedbackLoop(readInput())
	if signal != 14897241 {
		t.Error("Maximum amplified signal of part 2 is not 14897241", signal)
	}
}
