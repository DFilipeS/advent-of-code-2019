package day6

import (
	"github.com/golang-collections/collections/set"
	"github.com/golang-collections/collections/stack"
	"io/ioutil"
	"math"
	"strings"
)

func readInput() []string {
	inputBytes, _ := ioutil.ReadFile("./input.txt")

	return strings.Split(strings.Trim(string(inputBytes), "\n"), "\n")
}

func buildDirectedGraph(orbitsList []string) map[string][]string {
	graph := make(map[string][]string)

	for _, elem := range orbitsList {
		orbitList := strings.Split(elem, ")")
		innerOrbit := orbitList[0]
		outerOrbit := orbitList[1]
		graph[innerOrbit] = append(graph[innerOrbit], outerOrbit)
		graph[outerOrbit] = append(graph[outerOrbit])
	}

	return graph
}

func buildUndirectedGraph(orbitsList []string) map[string][]string {
	graph := make(map[string][]string)

	for _, elem := range orbitsList {
		orbitList := strings.Split(elem, ")")
		innerOrbit := orbitList[0]
		outerOrbit := orbitList[1]
		graph[innerOrbit] = append(graph[innerOrbit], outerOrbit)
		graph[outerOrbit] = append(graph[outerOrbit], innerOrbit)
	}

	return graph
}

func getTotalOrbits(graph map[string][]string) int {
	var totalEdges int
	for vertex := range graph {
		totalEdges += depthFirstSearch(graph, vertex) - 1
	}

	return totalEdges
}

func depthFirstSearch(graph map[string][]string, vertex string) int {
	visitedVertices := set.New()
	vertexStack := stack.New()

	var totalVertices int

	vertexStack.Push(vertex)
	for vertexStack.Len() > 0 {
		v := vertexStack.Pop()

		if !visitedVertices.Has(v) {
			visitedVertices.Insert(v)
			totalVertices++

			for _, adjacentVertex := range graph[v.(string)] {
				vertexStack.Push(adjacentVertex)
			}
		}
	}

	return totalVertices
}

func dijsktra(graph map[string][]string, sourceVertex string) map[string]int {
	unvisitedVertices := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)

	for vertex := range graph {
		dist[vertex] = math.MaxInt32
		unvisitedVertices[vertex] = true
	}
	dist[sourceVertex] = 0

	for len(unvisitedVertices) > 0 {
		var minimumDistanceVertex string
		var minimumDistance int = math.MaxInt32
		for vertex, distance := range dist {
			if unvisitedVertices[vertex] && distance < minimumDistance {
				minimumDistance = distance
				minimumDistanceVertex = vertex
			}
		}

		delete(unvisitedVertices, minimumDistanceVertex)

		for _, adjacentVertex := range graph[minimumDistanceVertex] {
			if unvisitedVertices[adjacentVertex] {
				alt := dist[minimumDistanceVertex] + 1
				if alt < dist[adjacentVertex] {
					dist[adjacentVertex] = alt
					prev[adjacentVertex] = minimumDistanceVertex
				}
			}
		}
	}

	return dist
}
