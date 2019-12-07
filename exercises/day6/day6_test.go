package day6

import (
	"testing"
)

func TestRunProgram(t *testing.T) {
	graph := buildDirectedGraph([]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"})
	totalOrbits := getTotalOrbits(graph)

	if totalOrbits != 42 {
		t.Error("Total orbits is not 42", totalOrbits)
	}

	graph = buildUndirectedGraph([]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"})
	minimumOrbitalTransfers := dijsktra(graph, "YOU")["SAN"] - 2
	if minimumOrbitalTransfers != 4 {
		t.Error("Minimum number of orbital transfers required is not 4", minimumOrbitalTransfers)
	}
}

func TestSolvePart1Input(t *testing.T) {
	graph := buildDirectedGraph(readInput())
	totalOrbits := getTotalOrbits(graph)

	if totalOrbits != 142915 {
		t.Error("Total orbits is not 142915", totalOrbits)
	}
}

func TestSolvePart2Input(t *testing.T) {
	graph := buildUndirectedGraph(readInput())
	minimumOrbitalTransfers := dijsktra(graph, "YOU")["SAN"] - 2

	if minimumOrbitalTransfers != 283 {
		t.Error("Minimum number of orbital transfers required is not 283", minimumOrbitalTransfers)
	}
}
