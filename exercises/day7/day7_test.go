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
